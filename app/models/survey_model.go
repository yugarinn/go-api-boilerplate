package models

import (
	"gorm.io/gorm"
	"time"
)

type Survey struct {
	gorm.Model
	ID uint
	Name string
	IsActive bool
	StartsAt time.Time
	EndsAt time.Time
	SurveySessions []SurveySession `gorm:"foreignKey:SessionID"`
}

func (Survey) TableName() string {
	return "surveys_surveys"
}
