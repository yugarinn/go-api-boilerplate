package controllers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/sergiouve/go-api-boilerplate/app/managers"
)

func ListSurveys(context *gin.Context) {
	surveys := managers.GetAllSurveys()

	context.JSON(http.StatusOK, gin.H{ "data": surveys })
}

func GetSurvey(context *gin.Context) {
	surveyId, _ := strconv.ParseInt(context.Param("survey"), 0, 64)
	survey := managers.GetSurvey(surveyId)

	context.JSON(http.StatusOK, gin.H{ "data": survey })
}
