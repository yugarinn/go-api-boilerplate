package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sergiouve/go-api-boilerplate/http/controllers"
)

func Register(router *gin.Engine) {
	router.GET("/users", controllers.GetUsers)
	router.GET("/status", controllers.GetStatus)
}
