package models

import (
	"github.com/jinzhu/gorm"
)

type SurveySession struct {
	gorm.Model
	ID uint
	Hash string
	StepIndex uint
	SessionID uint
}


func (SurveySession) TableName() string {
	return "surveys_sessions"
}
