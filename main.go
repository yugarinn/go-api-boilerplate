package main

import (
	"github.com/gin-gonic/gin"

	"github.com/yugarinn/go-api-boilerplate/http"
	"github.com/yugarinn/go-api-boilerplate/core"

)


func setupRouter(app *core.App) *gin.Engine {
	router := gin.Default()
	routes.Register(app, router)

	return router
}

func main() {
	app := core.BootstrapApplication()
	router := setupRouter(app)

	router.Run()
}
