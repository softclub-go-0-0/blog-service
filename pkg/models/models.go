package models

import (
	"gorm.io/gorm"
	"time"
)

type ID struct {
	ID uint `gorm:"primarykey"`
}

type Timestamps struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// ORM - Object-Relational Mapping

type User struct {
	ID        uint `gorm:"primarykey"`
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Posts    []Post
	Comments []Comment
}

type Post struct {
	ID
	UserID  uint
	Cover   string
	Title   string
	Content string
	Timestamps

	Comments []Comment
}

type Comment struct {
	gorm.Model
	PostID uint
	UserID uint
	Text   string
}
