package data

import (
	"errors"
	"log"

	"github.com/garindradeksa/socialmedia-mini/features/content"

	"gorm.io/gorm"
)

type contentData struct {
	db *gorm.DB
}

func New(db *gorm.DB) content.ContentData {
	return &contentData{
		db: db,
	}
}

func (cd *contentData) Add(userID uint, newContent content.Core) (content.Core, error) {
	cnv := CoreToData(newContent)
	cnv.UserID = uint(userID)
	err := cd.db.Create(&cnv).Error
	if err != nil {
		return content.Core{}, err
	}

	newContent.ID = cnv.ID

	return newContent, nil
}

func (cd *contentData) ContentDetail(contentID uint) (interface{}, error) {
	resContent := map[string]interface{}{}
	if err := cd.db.Raw("SELECT contents.id, users.avatar, users.username as username, contents.image, contents.caption, contents.created_at FROM contents JOIN users ON users.id = contents.user_id WHERE contents.id = ?", contentID).Find(&resContent).Error; err != nil {
		log.Println("Get content by content ID query error : ", err.Error())
		return nil, err
	}

	resComment := []map[string]interface{}{}

	if err := cd.db.Raw("SELECT comments.id, users.avatar, users.username, comments.text, comments.created_at FROM comments JOIN users ON users.id = comments.user_id WHERE comments.content_id = ?", contentID).Find(&resComment).Error; err != nil {
		log.Println("Get comments by content ID query error : ", err.Error())
		return nil, err
	}

	resContent["comments"] = resComment

	return resContent, nil
}

func (cd *contentData) ContentList() ([]content.Core, error) {
	res := []AllContents{}
	if err := cd.db.Table("contents").Joins("JOIN users ON users.id = contents.user_id").Select("contents.id, users.avatar as avatar, users.username as username, contents.image, contents.caption, contents.created_at as CreatedAt").Find(&res).Error; err != nil {
		log.Println("get all content query error : ", err.Error())
		return []content.Core{}, err
	}
	return AllListToCore(res), nil
}

func (cd *contentData) Update(userID uint, contentID uint, updatedContent content.Core) (content.Core, error) {
	getID := Contents{}
	err := cd.db.Where("id = ?", contentID).First(&getID).Error

	if err != nil {
		log.Println("get content error : ", err.Error())
		return content.Core{}, err
	}

	if getID.UserID != userID {
		log.Println("Unauthorized request")
		return content.Core{}, errors.New("Unauthorized request")
	}

	cnv := CoreToData(updatedContent)
	qry := cd.db.Where("id = ?", contentID).Updates(&cnv)
	if qry.RowsAffected <= 0 {
		log.Println("update content query error : data not found")
		return content.Core{}, errors.New("not found")
	}

	if err := qry.Error; err != nil {
		log.Println("update content query error : ", err.Error())
	}
	return updatedContent, nil
}

func (cd *contentData) Delete(userID uint, contentID uint) error {
	getID := Contents{}
	err := cd.db.Where("id = ?", contentID).First(&getID).Error

	if err != nil {
		log.Println("get content error : ", err.Error())
		return errors.New("failed to get content data")
	}

	if getID.UserID != userID {
		log.Println("unauthorized request")
		return errors.New("inauthorized request")
	}

	qryDelete := cd.db.Delete(&Contents{}, contentID)

	affRow := qryDelete.RowsAffected

	if affRow <= 0 {
		log.Println("No rows affected")
		return errors.New("failed to delete user content, data not found")
	}

	return nil
}
