package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aswinudhayakumar/golang-boilerplate/controller"
	"github.com/aswinudhayakumar/golang-boilerplate/model"
	"github.com/gorilla/mux"
)

//HandleUserRoutes handles all user routes
func HandleUserRoutes(userController controller.UserController, router *mux.Router) {

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		out := controller.SayHello(userController.DB)
		json.NewEncoder(w).Encode(out)
	}).Methods("GET")

	router.HandleFunc("/user/signup", func(w http.ResponseWriter, r *http.Request) {

		decoder := json.NewDecoder(r.Body)
		var out model.Output
		var t model.User
		err := decoder.Decode(&t)

		if err == nil {
			out = controller.Signup(userController.DB, t.Name, t.Email, t.Password)
			json.NewEncoder(w).Encode(out)
		} else {
			out.Code = 400
			out.Error = err.Error()
			json.NewEncoder(w).Encode(out)
		}
	}).Methods("POST")

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		out := controller.GetAllUsers(userController.DB)
		json.NewEncoder(w).Encode(out)
	}).Methods("GET")

	router.HandleFunc("/user/{id}/posts", func(w http.ResponseWriter, r *http.Request) {

		var out model.Output
		vars := mux.Vars(r)
		userID, err := strconv.Atoi(vars["id"])
		if err == nil {
			out = controller.GetPostsByUser(userController.DB, userID)
			json.NewEncoder(w).Encode(out)
		} else {
			out.Code = 400
			out.Error = err.Error()
			json.NewEncoder(w).Encode(out)
		}
	}).Methods("GET")

}
