package services

import (
	"fmt"

	core "github.com/yugarinn/go-api-boilerplate/core"
	inputs "github.com/yugarinn/go-api-boilerplate/app/users/inputs"
	managers "github.com/yugarinn/go-api-boilerplate/app/users/managers"
	models "github.com/yugarinn/go-api-boilerplate/app/users/models"
	utils "github.com/yugarinn/go-api-boilerplate/utils"
)


type CreateUserResult struct {
	User models.User
	Error error
}

type CreateUserValidationResult struct {
	UserValidation models.UserValidation
	Error error
}

const BOILERPLATE_PHONE_NUMBER = "+14178053542"

func CreateUserAndSendValidationCode(app *core.App, input inputs.CreateUserInput) CreateUserResult {
	user, creationError := managers.CreateUser(input)

	if creationError == nil {
		if utils.IsProduction() {
			go SendValidationSMS(app, user)
		} else {
			SendValidationSMS(app, user)
		}
	}

	return CreateUserResult{User: user, Error: creationError}
}

func SendValidationSMS(app *core.App, user models.User) CreateUserValidationResult {
	userValidation, error := managers.CreateValidationCodeFor(user.ID)

	if error == nil {
		toPhoneNumber := fmt.Sprintf("%s%s", user.PhonePrefix, user.PhoneNumber)
		fromPhoneNumber := BOILERPLATE_PHONE_NUMBER

		error = app.TwilioClient.SendSMS(toPhoneNumber, fromPhoneNumber, userValidation.Code)

		// TODO: logger
		fmt.Println(error)
	}

	return CreateUserValidationResult{UserValidation: userValidation, Error: error}
}
