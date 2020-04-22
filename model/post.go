package model

import "github.com/jinzhu/gorm"

//Post is a table to store all user posts
type Post struct {
	gorm.Model

	Title    string
	Body     string
	Likes    []Like    `gorm:"ForeignKey:PostID"`
	Comments []Comment `gorm:"ForeignKey:PostID"`
	AuthorID int       `gorm:"column:user_id"`
	Author   User
}

//PostOutput is a model
type PostOutput struct {
	ID            uint
	Title         string
	Body          string
	NumberOfLikes int
	User          interface{}
	Likes         interface{}
	Comments      interface{}
}

//Like is a table to store like of a post
type Like struct {
	gorm.Model

	PostID   int `gorm:"column:post_id"`
	Post     Post
	AuthorID int `gorm:"column:user_id"`
	Author   User
}

//LikeOutput isa struct
type LikeOutput struct {
	ID     uint
	Author interface{}
}

//Comment is a model
type Comment struct {
	gorm.Model

	Text     string
	PostID   int `gorm:"column:post_id"`
	Post     Post
	AuthorID int `gorm:"column:user_id"`
	Author   User
}

//CommentOutput is a model
type CommentOutput struct {
	ID     uint
	Text   string
	Author interface{}
}
