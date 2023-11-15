package data

import (
	"errors"
	"log"

	"github.com/garindradeksa/socialmedia-mini/features/comment"

	"gorm.io/gorm"
)

type commentData struct {
	db *gorm.DB
}

func New(db *gorm.DB) comment.CommentData {
	return &commentData{
		db: db,
	}
}

func (cd *commentData) Add(userID uint, newComment comment.Core, contentID uint) (comment.Core, error) {
	cnv := CoreToData(newComment)
	cnv.UserID = uint(userID)
	cnv.ContentID = contentID
	err := cd.db.Create(&cnv).Error
	if err != nil {
		return comment.Core{}, err
	}

	newComment.ID = cnv.ID
	newComment.ContentID = contentID

	return newComment, nil
}

func (cd *commentData) Delete(userID uint, commentID uint) error {
	getID := Comments{}
	err := cd.db.Where("id = ?", commentID).First(&getID).Error

	if err != nil {
		log.Println("Get comment error : ", err.Error())
		return errors.New("Failed to get comment data")
	}

	if getID.UserID != userID {
		log.Println("Unauthorized request")
		return errors.New("Unauthorized request")
	}

	qryDelete := cd.db.Delete(&Comments{}, commentID)

	affRow := qryDelete.RowsAffected

	if affRow <= 0 {
		log.Println("No rows affected")
		return errors.New("Failed to delete, data not found")
	}

	return nil
}

func (cd *commentData) CommentList() ([]comment.Core, error) {
	res := []AllComments{}
	if err := cd.db.Table("contents").Joins("JOIN users ON users.id = contents.user_id").Select("contents.id, contents.avatar, contents.username, contents.image, contents.caption, users.username as username").Find(&res).Error; err != nil {
		log.Println("get all content query error : ", err.Error())
		return []comment.Core{}, err
	}
	return AllListToCore(res), nil
}

// Done
