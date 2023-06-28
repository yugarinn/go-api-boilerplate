package users

import (
	"time"

	"gorm.io/gorm"
)


type UserValidation struct {
	gorm.Model
	ID uint64
	UserId uint64
	Code string
	IsUsed bool
	ExpiresAt time.Time
}

func (UserValidation) TableName() string {
	return "users_validations"
}
