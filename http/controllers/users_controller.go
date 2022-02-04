package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetUsers(context *gin.Context) {
	context.JSON(http.StatusOK, "OK")
}
