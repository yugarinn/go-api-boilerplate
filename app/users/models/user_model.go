package users

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID uint
	Email string
}

func (User) TableName() string {
	return "users_users"
}
