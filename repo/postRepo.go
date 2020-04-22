package repo

import (
	"github.com/aswinudhayakumar/golang-boilerplate/model"
	"github.com/jinzhu/gorm"
)

//AllPosts is a function to return all users
func AllPosts(db *gorm.DB) ([]model.PostOutput, string) {

	var posts []model.Post
	err := db.Preload("Author").Preload("Comments").Preload("Comments.Author").Preload("Likes").Preload("Likes.Author").Find(&posts).Error
	user := true
	comment := 1
	result := postsToPostOutput(posts, user, comment)

	if err == nil {
		return result, ""
	}
	return result, err.Error()
}

//AddNewPost is a function to add new post
func AddNewPost(db *gorm.DB, post model.Post) (string, string) {

	var user model.User
	notFound := db.Where("id = ?", post.AuthorID).Find(&user).RecordNotFound()

	if notFound == false {
		if err := db.Create(&post).Error; err != nil {
			return "", err.Error()
		}
		return "Post added successfully", ""
	}
	return "", "User not found for the given UserID"
}

//UpdatePost is a function to update post
func UpdatePost(db *gorm.DB, title string, body string, postID uint) (string, string) {

	var oldPost model.Post
	notFound := db.Where("id = ?", postID).Find(&oldPost).RecordNotFound()

	if notFound == false {
		if title != oldPost.Title {
			oldPost.Title = title
		}
		if body != oldPost.Body {
			oldPost.Body = body
		}
		if err := db.Save(&oldPost).Error; err != nil {
			return "Failed to update", err.Error()
		}
		return "Post successfully updated", ""
	}

	return "Failed to update", "Post not found"
}

//GetAllPostsOfUser is afunction to return all posts of a single user
func GetAllPostsOfUser(db *gorm.DB, userID int) ([]model.PostOutput, string) {

	var posts []model.Post
	err := db.Where("user_id = ?", userID).Preload("Comments").Preload("Comments.Author").Preload("Likes").Preload("Likes.Author").Find(&posts).Error
	user := false
	comment := 1
	result := postsToPostOutput(posts, user, comment)
	if err == nil {
		return result, ""
	}
	return result, err.Error()
}

//GetPostByID is a function to get single post by id
func GetPostByID(db *gorm.DB, postID int) (interface{}, string) {

	var post []model.Post
	var result []model.PostOutput
	err := db.Where("id = ?", postID).Preload("Comments").Preload("Likes").Preload("Comments.Author").Preload("Likes.Author").Find(&post).Error
	if err == nil {
		user := false
		comment := 2
		result = postsToPostOutput(post, user, comment)
		if len(result) > 0 {
			return result[0], ""
		}
		return "", "Post not found"
	}
	return result[0], err.Error()
}

//AddNewComment is a function to add new comment
func AddNewComment(db *gorm.DB, comment model.Comment) (string, string) {

	userNotFound := db.Where("id = ?", comment.AuthorID).RecordNotFound()
	postNotFound := db.Where("id = ?", comment.PostID).RecordNotFound()

	if userNotFound == false && postNotFound == false {
		if err := db.Save(&comment).Error; err != nil {
			return "", err.Error()
		}
		return "Comment added successfully", ""
	}

	return "", "User or post doesn't exists"
}

//DeletePostByID is a function to soft delete a post
func DeletePostByID(db *gorm.DB, postID int) (string, string) {

	var post model.Post
	err := db.Where("id = ?", postID).Find(&post).Error
	if err == nil {
		if err = db.Delete(&post).Error; err == nil {
			err = db.Where("post_id = ?", postID).Delete(model.Comment{}).Error
			if err == nil {
				return "Post deleted successfully", ""
			}
			return "", err.Error()
		}
		return "Post delete failed", err.Error()
	}
	return "Post not found ro delete", err.Error()
}

//DeleteCommentByID is a function to delete a single comment using id
func DeleteCommentByID(db *gorm.DB, commentID int) (string, string) {

	var comment model.Comment
	err := db.Where("id = ?", commentID).Find(&comment).Error
	if err == nil {
		if err = db.Delete(&comment).Error; err == nil {
			return "Comment deleted successfully", ""
		}
		return "", err.Error()
	}
	return "", err.Error()
}

//MutateLike is a function to add like
func MutateLike(db *gorm.DB, postID int, userID int, likeMethod string) (string, string) {

	var like model.Like
	var post model.Post
	var user model.User
	postNotFound := db.Where("id = ?", postID).Find(&post).RecordNotFound()
	userNotFound := db.Where("id = ?", userID).Find(&user).RecordNotFound()
	if postNotFound == false && userNotFound == false {
		userNotLikedPost := db.Where("post_id = ? AND user_id = ?", postID, userID).Find(&like).RecordNotFound()
		if userNotLikedPost == true && likeMethod == "LIKE" {

			var like model.Like
			like.AuthorID = userID
			like.PostID = postID
			err := db.Create(&like).Error
			if err == nil {
				return "Post liked successfully", ""
			}
			return "", err.Error()
		} else if userNotLikedPost == false && likeMethod == "UNLIKE" {

			var like model.Like
			like.AuthorID = userID
			like.PostID = postID
			err := db.Delete(&like).Error
			if err == nil {
				return "Post Disliked successfully", ""
			}
			return "", err.Error()
		}
		return "", "Post already liked by the user"
	}
	return "", "post or user not found"
}

func postsToPostOutput(posts []model.Post, user bool, comment int) []model.PostOutput {
	var result []model.PostOutput
	for _, elem := range posts {
		var r model.PostOutput
		r.ID = elem.ID
		r.Body = elem.Body
		r.Title = elem.Title
		if user == true {
			var userOut model.UserOutput
			userOut.ID = elem.Author.ID
			userOut.Email = elem.Author.Email
			userOut.Name = elem.Author.Email
			r.User = userOut
		}
		if comment > 0 {
			// r.Comment = elem.Comments
			var c model.CommentOutput
			var comments []model.CommentOutput
			var u model.UserOutput
			for _, commentData := range elem.Comments {
				c.ID = commentData.ID
				c.Text = commentData.Text
				u.Email = commentData.Author.Email
				u.ID = commentData.Author.ID
				u.Name = commentData.Author.Name
				c.Author = u
				comments = append(comments, c)
				if comment == 1 {
					break
				}
			}
			r.Comments = comments
		}
		var likes []model.LikeOutput
		for _, l := range elem.Likes {
			var like model.LikeOutput
			var u model.UserOutput
			u.Email = l.Author.Email
			u.ID = l.Author.ID
			u.Name = l.Author.Name
			like.ID = l.ID
			like.Author = u
			likes = append(likes, like)
		}
		r.Likes = likes
		r.NumberOfLikes = len(likes)
		result = append(result, r)
	}
	return result
}
