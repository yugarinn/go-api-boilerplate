package users

import (
	"gorm.io/gorm"
	"github.com/sergiouve/go-api-boilerplate/connections"
	"github.com/sergiouve/go-api-boilerplate/app/users/models"
)


var database *gorm.DB = connections.Database()

func GetAllUsers() []users.User {
	var users []users.User
	database.Find(&users)

	return users
}

func GetUser(id int64) users.User {
	var user users.User
	database.Where("id", id).First(&user)

	return user
}
