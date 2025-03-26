package models

import "time"

type UserAnswer struct {
	UserID     uint      `gorm:"primaryKey;not null"`
	SessionID  uint      `gorm:"primaryKey;not null"`
	QuestionID uint      `gorm:"primaryKey;not null"`
	AnswerID   uint      `gorm:"not null"`
	User       *User     `gorm:"foreignKey:UserID;references:ID"`
	Session    *Session  `gorm:"foreignKey:SessionID;references:ID"`
	Question   *Question `gorm:"foreignKey:QuestionID;references:ID"`
	Answer     *Answer   `gorm:"foreignKey:AnswerID;references:ID"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
