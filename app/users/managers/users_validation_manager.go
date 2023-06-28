package users

import (
	"math/rand"
	"time"

	users "github.com/yugarinn/go-api-boilerplate/app/users/models"
)


var VALIDATION_CODE_CHARSET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
var VALIDATION_CODE_LENGTH = 6
var VALIDATION_CODE_LIFETIME_IN_SECONDS = 180

func CreateValidationCodeFor(userId uint64) (users.UserValidation, error) {
    validation := users.UserValidation{UserId: userId, Code: generateValidationCode(), IsUsed: false, ExpiresAt: generateExpirationDate()}
	result := database.Create(&validation)

	return validation, result.Error
}

func GetUserValidation(validationId uint64) (users.UserValidation, error) {
	var validation users.UserValidation
	result := database.Where("id", validationId).First(&validation)

	return validation, result.Error
}

func SetValidationAsUsed(id uint64) (users.UserValidation, error) {
	validation, retrievalError := GetUserValidation(id)

	if retrievalError != nil {
		return users.UserValidation{}, retrievalError
	}

	validation.IsUsed = true
	updateResult := database.Save(&validation)

	return validation, updateResult.Error
}

func generateValidationCode() string {
    rand.Seed(time.Now().UnixNano())
    code := make([]byte, VALIDATION_CODE_LENGTH)

    for i := range code {
        code[i] = VALIDATION_CODE_CHARSET[rand.Intn(len(VALIDATION_CODE_CHARSET))]
    }

	var codeAlreadyExists int64
	database.Model(&users.UserValidation{}).Where("code = ?", code).Count(&codeAlreadyExists)

	if codeAlreadyExists != 0 {
		return generateValidationCode()
	}

    return string(code)
}

func generateExpirationDate() time.Time {
	now := time.Now()

	return now.Add(time.Duration(VALIDATION_CODE_LIFETIME_IN_SECONDS) * time.Second).UTC()
}
