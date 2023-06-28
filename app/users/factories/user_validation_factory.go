package factories

import (
	"time"

	"github.com/yugarinn/go-api-boilerplate/app/users/models"
)


type UserValidationFactoryInput struct {
	UserID uint64
	Code string
	IsUsed bool
	ExpiresAt time.Time
}

func CreateUserValidation(input UserValidationFactoryInput) users.UserValidation {
    validation := users.UserValidation{UserId: input.UserID, Code: input.Code, IsUsed: input.IsUsed, ExpiresAt: input.ExpiresAt}

	if len(validation.Code) == 0 {
		validation.Code = "123ABC"
	}

	database.Create(&validation)

	return validation
}

func CreateUserValidationList(number int, properties UserValidationFactoryInput) []users.UserValidation {
	var validations []users.UserValidation

	for i := 0; i < number; i++ {
		validations = append(validations, CreateUserValidation(properties))
	}

	return validations
}
