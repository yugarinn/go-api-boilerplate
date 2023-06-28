package middlewares

import (
	"errors"
	"fmt"
	"strconv"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	models "github.com/yugarinn/go-api-boilerplate/app/users/models"
	managers "github.com/yugarinn/go-api-boilerplate/app/users/managers"
)

func CheckAccessToken(context *gin.Context) {
    providedToken := context.GetHeader("Access-Token")

    if false && providedToken == "" {
        context.AbortWithError(401, errors.New("unauthorized"))
        return
    } else {
		user, validationError := validateToken(providedToken)

        if validationError != nil {
            context.AbortWithError(401, errors.New("unauthorized"))
            return
        }

		context.Set("user", user)
		context.Next()
	}
}

func validateToken(providedToken string) (models.User, error) {
	secret := os.Getenv("JWT_SECRET_KEY")

    token, err := jwt.Parse(providedToken, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }

        return []byte(secret), nil
    })

    if err != nil {
        return models.User{}, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        userIDStr, ok := claims["userId"].(string)

        if !ok {
            return models.User{}, fmt.Errorf("error getting ID from token")
        }

		userID, _ := strconv.ParseUint(userIDStr, 10, 64)
        user, userRetrievalError := managers.GetUser(userID)

        if userRetrievalError != nil {
            return models.User{}, err
        }

        return user, nil
    } else {
        return models.User{}, fmt.Errorf("invalid token")
    }
}
