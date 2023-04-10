package models

import "database/sql"

// ORM - Object-Relational Mapping

type User struct {
	ID        uint
	Name      string
	Email     string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}
