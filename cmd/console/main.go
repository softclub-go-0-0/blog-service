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

		fmt.Println("\t7. Update a user")
		fmt.Println("\t8. Delete a user")

		var command int
		fmt.Print("\nYour command: ")
		fmt.Scan(&command)

		switch command {
		case 0:
			fmt.Println("Bye!")
			return
		case 1:
			fmt.Println("\nCreating a user...")
			var user models.User
			fmt.Print("Enter user's name: ")
			fmt.Scan(&user.Name)
			fmt.Print("Enter user's email: ")
			fmt.Scan(&user.Email)
			db.Create(&user)
			fmt.Println("Successfully created new user with ID:", user.ID)
		case 2:
			fmt.Println("\nThe list of users in DB:")
			var users []models.User
			db.Preload("Posts").Preload("Comments").Find(&users)
			for _, user := range users {
				fmt.Printf("%d. %s\t%s\t%d\t%d\n", user.ID, user.Name, user.Email, len(user.Posts), len(user.Comments))
			}
		case 4:
			fmt.Println("\nThe list of posts in DB:")
			var posts []models.Post
			db.Preload("Comments").Preload("Author").Find(&posts)
			for _, post := range posts {
				fmt.Printf("%d. %s\t%s\t%s\t%s\t%d\n", post.ID.ID, post.Cover, post.Title, post.Content, post.Author.Name, len(post.Comments))
			}
		case 7:
			fmt.Print("Give me an ID of existing user: ")
			var id uint
			fmt.Scan(&id)

			var user models.User

			fmt.Println("Empty user:", user)

			result := db.First(&user, id)
			if result.Error != nil {
				log.Fatal(result.Error)
			}

			fmt.Println("Here is your user:", user)

			fmt.Scan(&user.Name)
			fmt.Scan(&user.Email)

			db.Save(&user)

			fmt.Println("User has been updated")
			fmt.Println(user)
		}
	}
}
