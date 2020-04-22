package controller

import (
	"github.com/aswinudhayakumar/golang-boilerplate/model"
	"github.com/aswinudhayakumar/golang-boilerplate/repo"
	"github.com/jinzhu/gorm"
)

//PostController is a struct for database
type PostController struct {
	DB *gorm.DB
}

//GetAllPosts is a function to return all posts from database
func GetAllPosts(db *gorm.DB) model.Output {
	var out model.Output
	out.Code = 200
	out.Data, out.Error = repo.AllPosts(db)
	if out.Error != "" {
		out.Code = 400
	}
	return out
}

//AddNewPost is a function to connect addnewpost func in repo
func AddNewPost(db *gorm.DB, title string, body string, userid int) model.Output {

	var out model.Output
	var post model.Post

	out.Code = 200
	post.Title = title
	post.Body = body
	post.AuthorID = userid

	out.Data, out.Error = repo.AddNewPost(db, post)
	if out.Error != "" {
		out.Code = 400
	}
	return out
}

//UpdatePost is a function to update post
func UpdatePost(db *gorm.DB, title string, body string, postID uint) model.Output {

	var out model.Output
	out.Data, out.Error = repo.UpdatePost(db, title, body, postID)
	if out.Error != "" {
		out.Code = 400
	}
	return out
}

//GetPostsByUser is a function
func GetPostsByUser(db *gorm.DB, UserID int) model.Output {

	var out model.Output
	out.Code = 200

	out.Data, out.Error = repo.GetAllPostsOfUser(db, UserID)
	if out.Error != "" {
		out.Code = 400
	}
	return out
}

//AddNewComment is to add new comment
func AddNewComment(db *gorm.DB, text string, postID int, authorID int) model.Output {

	var comment model.Comment
	var out model.Output
	out.Code = 200
	comment.Text = text
	comment.PostID = postID
	comment.AuthorID = authorID

	out.Data, out.Error = repo.AddNewComment(db, comment)
	if out.Error != "" {
		out.Code = 400
	}
	return out
}

//GetPostByID is a function to get a single post
func GetPostByID(db *gorm.DB, postID int) model.Output {
	var out model.Output
	out.Code = 200
	out.Data, out.Error = repo.GetPostByID(db, postID)

	if out.Error != "" {
		out.Code = 400
	}
	return out
}

//DeletePostByID is a function to soft delete apost
func DeletePostByID(db *gorm.DB, postID int) model.Output {

	var out model.Output
	out.Code = 200
	out.Data, out.Error = repo.DeletePostByID(db, postID)
	if out.Error != "" {
		out.Code = 400
	}
	return out
}

//DeleteCommentByID is a function to delete a single comment by ID
func DeleteCommentByID(db *gorm.DB, commentID int) model.Output {

	var out model.Output
	out.Code = 200
	out.Data, out.Error = repo.DeleteCommentByID(db, commentID)
	if out.Error != "" {
		out.Code = 400
	}
	return out
}

//AddLike is a function to add like for a post
func AddLike(db *gorm.DB, postID int, userID int) model.Output {

	var out model.Output
	out.Code = 200
	out.Data, out.Error = repo.MutateLike(db, postID, userID, "LIKE")
	if out.Error != "" {
		out.Code = 400
	}
	return out
}

//UnLike is a function to add like for a post
func UnLike(db *gorm.DB, postID int, userID int) model.Output {

	var out model.Output
	out.Code = 200
	out.Data, out.Error = repo.MutateLike(db, postID, userID, "UNLIKE")
	if out.Error != "" {
		out.Code = 400
	}
	return out
}
