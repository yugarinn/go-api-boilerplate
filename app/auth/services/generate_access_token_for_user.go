package services

import (
	authModels "github.com/yugarinn/go-api-boilerplate/app/auth/models"
	authManagers "github.com/yugarinn/go-api-boilerplate/app/auth/managers"
	usersManagers "github.com/yugarinn/go-api-boilerplate/app/users/managers"
)


type GenerateAccessTokenForUserResult struct {
	AccessToken authModels.AccessToken
	Error error
}

func GenerateAccessTokenForUser(userId uint64) GenerateAccessTokenForUserResult {
	user, userRetrievalError := usersManagers.GetUser(userId)

	if userRetrievalError != nil {
		return GenerateAccessTokenForUserResult{AccessToken: authModels.AccessToken{}, Error: userRetrievalError}
	}

	accessToken, jwtGenerationError := authManagers.GenerateAccessTokenForUser(user)

	if userRetrievalError != nil {
		return GenerateAccessTokenForUserResult{AccessToken: authModels.AccessToken{}, Error: jwtGenerationError}
	}

	return GenerateAccessTokenForUserResult{AccessToken: accessToken, Error: nil}
}
