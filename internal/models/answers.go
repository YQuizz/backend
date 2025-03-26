package models

import "time"

type Answer struct {
	ID         uint      `gorm:";primaryKey"`
	QuestionID uint      `gorm:"not null"`
	Question   *Question `gorm:"foreignKey:QuestionID;references:ID"`
	Text       string    `gorm:"not null"`
	IsCorrect  bool      `gorm:"not null;default:false"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
