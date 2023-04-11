package main

import (
	"fmt"
	"github.com/softclub-go-0-0/blog-service/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const (
	DBHost     = "localhost" // 127.0.0.1
	DBPort     = 5432
	DBUser     = "postgres"
	DBPassword = "postgres"
	DBName     = "blog_service"
)

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Dushanbe", DBHost, DBPort, DBUser, DBPassword, DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected")

	db.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Comment{},
	)

	for {
		fmt.Println("\nApplication's functional:")
		fmt.Println("\t0. Exit")

		fmt.Println("\t1. Create a user")
		fmt.Println("\t2. Get the list all users")

		fmt.Println("\t3. Create a post")
		fmt.Println("\t4. Get the list all posts")

		fmt.Println("\t5. Create a comment")
		fmt.Println("\t6. Get the list all comments")

		var command int
		fmt.Print("\nYour command: ")
		fmt.Scan(&command)

		switch command {
		case 0:
			fmt.Println("Bye!")
			return
		case 1:
			fmt.Println("Creating a user...")
			var user models.User
			fmt.Print("Enter user's name: ")
			fmt.Scan(&user.Name)
			fmt.Print("Enter user's email: ")
			fmt.Scan(&user.Email)
			db.Create(&user)
			fmt.Println("Successfully created new user with ID:", user.ID)
		}
	}
}
