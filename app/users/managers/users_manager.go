package users

import (
	"fmt"
	"errors"

	"github.com/yugarinn/go-api-boilerplate/app/users/inputs"
	"github.com/yugarinn/go-api-boilerplate/app/users/models"
)


func CreateUser(input inputs.CreateUserInput) (users.User, error) {
	if userPhoneNumberIsAlreadyRegistered(input.PhoneNumber) {
		return users.User{}, errors.New("phone_number_already_registered")
	}

    user := users.User{Email: input.Email, Name: input.Name, LastName: input.LastName, PhonePrefix: input.PhonePrefix, PhoneNumber: input.PhoneNumber, CountryCode: input.CountryCode}
	result := database.Create(&user)

	return user, result.Error
}

func GetPaginatedUsers(page int, size int) ([]users.User, error) {
	var users []users.User
	result := database.Find(&users)

	return users, result.Error
}

func GetUser(id uint64) (users.User, error) {
	var user users.User
	result := database.Where("id", id).First(&user)

	return user, result.Error
}

func GetUserBy(field string, value any) (users.User, error) {
	var user users.User
	result := database.Where(field, value).First(&user)

	return user, result.Error
}

func SetUserAsConfirmed(id uint64) (users.User, error) {
	user, retrievalError := GetUser(id)

	if retrievalError != nil {
		return users.User{}, retrievalError
	}

	user.IsConfirmed = true
	updateResult := database.Save(&user)

	return user, updateResult.Error
}

func userPhoneNumberIsAlreadyRegistered(email string) bool {
	var alreadyExists bool
	database.Raw(fmt.Sprintf("SELECT COUNT(1) FROM users_users WHERE phone_number='%s';", email)).Scan(&alreadyExists)

	return alreadyExists
}
