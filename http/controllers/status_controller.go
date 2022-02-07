package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetStatus(context *gin.Context) {
	context.Status(http.StatusOK)
}
