package auth

import (
	"time"

	"gorm.io/gorm"
)

type AccessToken struct {
	gorm.Model
	ID uint64
	PasswordId uint64
	UserId uint64
	Token string
	ExpiresAt time.Time
}

func (AccessToken) TableName() string {
	return "auth_access_tokens"
}
