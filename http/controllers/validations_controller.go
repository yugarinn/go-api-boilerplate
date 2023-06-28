package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	core "github.com/yugarinn/go-api-boilerplate/core"
	inputs "github.com/yugarinn/go-api-boilerplate/app/users/inputs"
	usersServices "github.com/yugarinn/go-api-boilerplate/app/users/services"
	authServices "github.com/yugarinn/go-api-boilerplate/app/auth/services"
	responses "github.com/yugarinn/go-api-boilerplate/http/responses"
)


func CreateUserValidation(app *core.App, context *gin.Context) {
	userId, _ := strconv.ParseUint(context.Param("userId"), 10, 64)
	input := inputs.GetUserInput{UserID: userId}
	validate := validator.New()

	validationErrors := validate.Struct(&input)

	if validationErrors != nil {
		FailWithHttpCode(context, 422, validationErrors.Error())
		return
	}

	getUserResult := usersServices.GetUser(input)
	result := usersServices.SendValidationSMS(app, getUserResult.User)

	context.JSON(http.StatusCreated, responses.SerializeValidation(result.UserValidation))
}

func ValidateUser(context *gin.Context) {
	userId, _ := strconv.ParseUint(context.Param("userId"), 10, 64)
	validationId, _ := strconv.ParseUint(context.Param("validationId"), 10, 64)
	input := inputs.ValidateUserInput{UserID: userId, ValidationID: validationId}
	validate := validator.New()

	context.BindJSON(&input)
	validationErrors := validate.Struct(input)

	if validationErrors != nil {
		FailWithHttpCode(context, 422, validationErrors.Error())
		return
	}

	validateUserResult := usersServices.ValidateUser(input)

	if validateUserResult.Success == false {
		FailWithHttpCode(context, 422, validateUserResult.Error.Error())
		return
	}

	result := authServices.GenerateAccessTokenForUser(userId);

	context.JSON(http.StatusCreated, responses.SerializeAccessToken(result.AccessToken))
}
