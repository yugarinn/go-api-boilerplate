package services

import (
	inputs "github.com/yugarinn/go-api-boilerplate/app/users/inputs"
	managers "github.com/yugarinn/go-api-boilerplate/app/users/managers"
	models "github.com/yugarinn/go-api-boilerplate/app/users/models"
)


type GetPaginatedUsersResult struct {
	Users []models.User
	Error error
}

func GetPaginatedUsers(input inputs.GetPaginatedUsersInput) GetPaginatedUsersResult {
	users, userRetrievalError := managers.GetPaginatedUsers(input.Page, input.Size)

	return GetPaginatedUsersResult{Users: users, Error: userRetrievalError}
}
