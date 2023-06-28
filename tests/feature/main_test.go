package tests

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/yugarinn/go-api-boilerplate/app/auth/models"
	"github.com/yugarinn/go-api-boilerplate/app/users/models"
	authServices "github.com/yugarinn/go-api-boilerplate/app/auth/services"
	"github.com/yugarinn/go-api-boilerplate/connections"
	"github.com/yugarinn/go-api-boilerplate/http"
	"github.com/yugarinn/go-api-boilerplate/core"
	"github.com/yugarinn/go-api-boilerplate/tests/mocks"
)


var database *gorm.DB = connections.Database()

func SetupRouter() (*core.App, *gin.Engine) {
	gin.SetMode("test")

	app := mockApp()
	router := gin.Default()
	routes.Register(app, router)

	return app, router
}

func ResetDatabase() {
	databaseTeardown()
	databaseSetup()
}

func AuthenticateAs(userId uint64, request *http.Request) {
	accessToken := authServices.GenerateAccessTokenForUser(userId)
	request.Header.Set("Access-Token", accessToken.AccessToken.Token)
}

func DatabaseHas(tableName string, whereClause string) bool {
	var count int64
	database.Raw(fmt.Sprintf("SELECT COUNT(1) FROM %s WHERE %s;", tableName, whereClause)).Scan(&count)

	return count > 0
}

func DatabaseHasCount(tableName string, expectedCount int, whereClause string) bool {
	var actualCount int
	database.Raw(fmt.Sprintf("SELECT COUNT(1) FROM %s WHERE %s;", tableName, whereClause)).Scan(&actualCount)

	return expectedCount == actualCount
}

func DatabaseMissing(tableName string, whereClause string) bool {
	return !DatabaseHas(tableName, whereClause)
}

func databaseSetup() {
	database.AutoMigrate(&auth.AccessToken{})
	database.AutoMigrate(&users.User{})
	database.AutoMigrate(&users.UserValidation{})
}

func databaseTeardown() {
	database.Migrator().DropTable(&auth.AccessToken{})
	database.Migrator().DropTable(&users.User{})
	database.Migrator().DropTable(&users.UserValidation{})
}

func mockApp() *core.App {
    app := &core.App{
        TwilioClient: &mocks.TwilioMock{},
    }

	return app
}
