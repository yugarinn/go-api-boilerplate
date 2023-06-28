package users

import (
	"time"

	"gorm.io/gorm"
)

type UserPassword struct {
	gorm.Model
	ID uint64
	UserId uint64
	Hash string
	IsUsed bool
	ExpiresAt time.Time
}

func (UserPassword) TableName() string {
	return "users_passwords"
}
