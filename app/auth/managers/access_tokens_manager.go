package auth

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"

	auth "github.com/yugarinn/go-api-boilerplate/app/auth/models"
	users "github.com/yugarinn/go-api-boilerplate/app/users/models"
	"github.com/yugarinn/go-api-boilerplate/connections"
)


var database *gorm.DB = connections.Database()

type JWTClaims struct {
    UserID    string `json:"userId"`
    UserEmail string `json:"userEmail"`
    jwt.StandardClaims
}

func GenerateAccessTokenForUser(user users.User) (auth.AccessToken, error) {
	expirationDate := generateExpirationDate()

	jwtClaims := JWTClaims{
        UserEmail: user.Email,
        UserID: strconv.FormatUint(user.ID, 10),
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
        },
	}

	secret := os.Getenv("JWT_SECRET_KEY")

	token, _ := generateJWT(jwtClaims, secret)
    accessToken := auth.AccessToken{UserId: user.ID, Token: token, ExpiresAt: expirationDate}
	result := database.Create(&accessToken)

	return accessToken, result.Error
}

func GetAccessTokenByBy(field string, value any) (auth.AccessToken, error) {
	var accessToken auth.AccessToken
	result := database.Where(field, value).First(&accessToken)

	return accessToken, result.Error
}

func generateExpirationDate() time.Time {
	return time.Now().Add(time.Hour * 24)
}


func generateJWT(claims JWTClaims, secret string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    return token.SignedString([]byte(secret))
}
