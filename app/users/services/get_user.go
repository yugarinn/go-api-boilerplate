package services

import (
	inputs "github.com/yugarinn/go-api-boilerplate/app/users/inputs"
	managers "github.com/yugarinn/go-api-boilerplate/app/users/managers"
	models "github.com/yugarinn/go-api-boilerplate/app/users/models"
)


type GetUserResult struct {
	User models.User
	Error error
}

func GetUser(input inputs.GetUserInput) GetUserResult {
	user, userRetrievalError := managers.GetUser(input.UserID)

	return GetUserResult{User: user, Error: userRetrievalError}
}
