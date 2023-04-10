package main

import (
	"fmt"
	"github.com/softclub-go-0-0/blog-service/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const (
	host     = "localhost" // 127.0.0.1
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "blog_service"
)

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Dushanbe", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("successfully connected")

	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})

	user := models.User{}
	//
	//db.Create(&user)

	db.First(&user, 1)

	fmt.Println("Selected user:", user)

	post := models.Post{
		UserID:  user.ID,
		Title:   "SOME TEXT",
		Content: "Dubai",
	}

	db.Create(&post)

	fmt.Println("Created post:", post)

	comment := models.Comment{
		PostID: post.ID.ID,
		UserID: user.ID,
		Text:   "Some comment",
	}

	db.Create(&comment)

	fmt.Println("Created comment:", comment)

}
