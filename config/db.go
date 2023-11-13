package config

import (
	"fmt"
	"log"

	comment "github.com/garindradeksa/socialmedia-mini/features/comment/data"
	content "github.com/garindradeksa/socialmedia-mini/features/content/data"
	user "github.com/garindradeksa/socialmedia-mini/features/user/data"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(ac AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		ac.DBUser, ac.DBPass, ac.DBHost, ac.DBPort, ac.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Database connection error : ", err.Error())
		return nil
	}

	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(user.Users{})
	db.AutoMigrate(content.Contents{})
	db.AutoMigrate(comment.Comments{})
}
