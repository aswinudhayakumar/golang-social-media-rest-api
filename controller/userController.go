package controller

import (
	"regexp"

	"github.com/aswinudhayakumar/golang-boilerplate/model"
	"github.com/aswinudhayakumar/golang-boilerplate/repo"
	"github.com/jinzhu/gorm"
)

//UserController is a struct for database
type UserController struct {
	DB *gorm.DB
}

//SayHello is a function to add new user
func SayHello(db *gorm.DB) model.Output {

	var out model.Output
	out.Data = "Hello from controller in"
	out.Code = 200
	out.Error = ""

	return out

}

//Signup is a function to create user
func Signup(db *gorm.DB, name string, email string, password string) model.Output {

	var out model.Output
	out.Code = 200

	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if emailRegex.MatchString(email) == false {
		out.Error = "Email validation failed"
		return out
	}

	var user model.User
	user.Name = name
	user.Email = email
	user.Password = password

	data, err := repo.Signup(db, user)

	if err != "" {
		out.Code = 400
	}
	out.Data = data
	out.Error = err

	return out
}

//GetAllUsers will return all users
func GetAllUsers(db *gorm.DB) model.Output {

	var out model.Output
	out.Code = 200
	out.Data, out.Error = repo.AllUsers(db)
	if out.Error != "" {
		out.Code = 400
	}
	return out
}
