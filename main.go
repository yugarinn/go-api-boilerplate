package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sergiouve/go-api-boilerplate/http"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	routes.Register(router)

	return router
}

func main() {
	router := setupRouter()
	router.Run()
}
