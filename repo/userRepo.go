package repo

import (
	"github.com/aswinudhayakumar/golang-boilerplate/model"
	"github.com/jinzhu/gorm"
)

//Signup is a function to create new user
func Signup(db *gorm.DB, user model.User) (string, string) {

	var res model.User
	data := db.Where("Email = ?", user.Email).First(&res).RecordNotFound()

	if data == true {
		if err := db.Create(&user).Error; err != nil {
			return "", err.Error()
		}
		return "User added successfully", ""
	}

	return "", "Email already exists - Error"
}

//AllUsers is a function to return all users
func AllUsers(db *gorm.DB) ([]model.UserOutput, string) {
	var users []model.User
	var result []model.UserOutput
	err := db.Preload("Posts").Find(&users).Error
	if err == nil {
		for _, elem := range users {
			var r model.UserOutput
			r.ID = elem.ID
			r.Name = elem.Name
			r.Email = elem.Email
			result = append(result, r)
		}
		return result, ""
	}
	return result, err.Error()
}
