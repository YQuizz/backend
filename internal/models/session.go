package models

import "time"

type Session struct {
	ID             uint                 `gorm:"primaryKey"`
	QuizID         uint                 `gorm:"not null"`
	StartTime      time.Time            `gorm:"not null"`
	EndTime        time.Time            `gorm:"not null"`
	Quiz           *Quiz                `gorm:"foreignKey:QuizID;references:ID"`
	Participants   []SessionParticipant `gorm:"foreignKey:SessionID;references:ID"`
	UserAnswers    []UserAnswer         `gorm:"foreignKey:SessionID;references:ID"`
	MonitoringLogs []MonitoringLog      `gorm:"foreignKey:SessionID;references:ID"`
}
