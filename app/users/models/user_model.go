package users

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID uint64
	Email string
	PhonePrefix string
	PhoneNumber string
	CountryCode string
	Name string
	LastName string
	IsConfirmed bool
}

func (User) TableName() string {
	return "users_users"
}
