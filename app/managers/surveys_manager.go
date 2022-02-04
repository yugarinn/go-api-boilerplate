package managers

import (
	"gorm.io/gorm"
	"github.com/sergiouve/go-api-boilerplate/connections"
	"github.com/sergiouve/go-api-boilerplate/app/models"
)


var database *gorm.DB = connections.Database()

func GetAllSurveys() []models.Survey {
	var surveys []models.Survey
	database.Find(&surveys)

	return surveys
}

func GetSurvey(id int64) models.Survey {
	var survey models.Survey
	database.Where("id", id).First(&survey)

	return survey
}
