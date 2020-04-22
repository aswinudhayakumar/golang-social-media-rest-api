package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aswinudhayakumar/golang-boilerplate/controller"
	"github.com/aswinudhayakumar/golang-boilerplate/model"
	"github.com/gorilla/mux"
)

//HandlePostsRoutes handles all user route
func HandlePostsRoutes(postContoller controller.PostController, router *mux.Router) {

	router.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		out := controller.GetAllPosts(postContoller.DB)
		json.NewEncoder(w).Encode(out)
	}).Methods("GET")

	router.HandleFunc("/post/{id}", func(w http.ResponseWriter, r *http.Request) {
		var out model.Output
		vars := mux.Vars(r)
		PostID, err := strconv.Atoi(vars["id"])

		if err == nil {
			out := controller.GetPostByID(postContoller.DB, PostID)
			json.NewEncoder(w).Encode(out)
		} else {
			out.Code = 400
			out.Error = err.Error()
			json.NewEncoder(w).Encode(out)
		}
	}).Methods("GET")

	router.HandleFunc("/post/add", func(w http.ResponseWriter, r *http.Request) {

		decoder := json.NewDecoder(r.Body)
		var post model.Post
		var out model.Output
		err := decoder.Decode(&post)

		if err == nil && post.AuthorID > 0 {
			out = controller.AddNewPost(postContoller.DB, post.Title, post.Body, post.AuthorID)
			json.NewEncoder(w).Encode(out)
		} else {
			out.Code = 400
			if err != nil {
				out.Error = err.Error()
			}
			out.Error = "AuthourID is invalid"
			json.NewEncoder(w).Encode(out)
		}
	}).Methods("POST")

	router.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var post model.Post
		var out model.Output
		err := decoder.Decode(&post)
		if err == nil && post.ID > 0 {
			out = controller.UpdatePost(postContoller.DB, post.Title, post.Body, post.ID)
			json.NewEncoder(w).Encode(out)
		} else {
			out.Code = 400
			out.Error = err.Error()
			json.NewEncoder(w).Encode(out)
		}
	}).Methods("PUT")

	router.HandleFunc("/post/{id}", func(w http.ResponseWriter, r *http.Request) {
		var out model.Output
		vars := mux.Vars(r)
		PostID, err := strconv.Atoi(vars["id"])

		if err == nil && PostID > 0 {
			out = controller.DeletePostByID(postContoller.DB, PostID)
			json.NewEncoder(w).Encode(out)
		} else {
			out.Code = 400
			out.Error = err.Error()
			json.NewEncoder(w).Encode(out)
		}
	}).Methods("DELETE")

	router.HandleFunc("/comment/{id}", func(w http.ResponseWriter, r *http.Request) {
		var out model.Output
		vars := mux.Vars(r)
		commentID, err := strconv.Atoi(vars["id"])

		if err == nil && commentID > 0 {
			out = controller.DeleteCommentByID(postContoller.DB, commentID)
			json.NewEncoder(w).Encode(out)
		} else {
			out.Code = 400
			out.Error = err.Error()
			json.NewEncoder(w).Encode(out)
		}
	}).Methods("DELETE")

	router.HandleFunc("/post/comment", func(w http.ResponseWriter, r *http.Request) {

		decoder := json.NewDecoder(r.Body)
		var comment model.Comment
		var out model.Output
		err := decoder.Decode(&comment)

		if err == nil && comment.AuthorID > 0 && comment.PostID > 0 {
			out = controller.AddNewComment(postContoller.DB, comment.Text, comment.PostID, comment.AuthorID)
			json.NewEncoder(w).Encode(out)
		} else {
			out.Code = 400
			if err != nil {
				out.Error = err.Error()
			}
			out.Error = "AuthourID and PostID is invalid"
			json.NewEncoder(w).Encode(out)
		}
	}).Methods("POST")

	router.HandleFunc("/user/post/like", func(w http.ResponseWriter, r *http.Request) {

		v := r.URL.Query()
		var out model.Output
		userID, userErr := strconv.Atoi(v.Get("userID"))
		postID, postErr := strconv.Atoi(v.Get("postID"))
		if userErr != nil || postErr != nil {
			out.Code = 400
			out.Error = "Error processing inputs"
			json.NewEncoder(w).Encode(out)
		} else {
			out = controller.AddLike(postContoller.DB, postID, userID)
			json.NewEncoder(w).Encode(out)
		}
	}).Methods("POST")

	router.HandleFunc("/user/post/unlike", func(w http.ResponseWriter, r *http.Request) {

		v := r.URL.Query()
		var out model.Output
		userID, userErr := strconv.Atoi(v.Get("userID"))
		postID, postErr := strconv.Atoi(v.Get("postID"))
		if userErr != nil || postErr != nil {
			out.Code = 400
			out.Error = "Error processing inputs"
			json.NewEncoder(w).Encode(out)
		} else {
			out = controller.UnLike(postContoller.DB, postID, userID)
			json.NewEncoder(w).Encode(out)
		}
	}).Methods("POST")

}
