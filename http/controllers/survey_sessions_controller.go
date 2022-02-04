package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func CreateSurveySession(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{ "data": "TODO" })
}
