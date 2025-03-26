package models

import "time"

type SessionParticipant struct {
	SessionID       uint      `gorm:"primaryKey"`
	ParticipantType string    `gorm:"type:user_role;default:'student'"` //student, teacher, admin
	ParticipantID   uint      `gorm:"primaryKey"`
	Session         *Session  `gorm:"foreignKey:SessionID;references:ID"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
}
