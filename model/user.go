package model

import "github.com/jinzhu/gorm"

//User model
type User struct {
	gorm.Model

	Name     string
	Email    string `gorm:"unique;not null"`
	Password string
	Posts    []Post    `gorm:"ForeignKey:AuthorID"`
	Comments []Comment `gorm:"ForeignKey:AuthorID"`
}

//UserOutput model
type UserOutput struct {
	ID    uint
	Name  string
	Email string
}
