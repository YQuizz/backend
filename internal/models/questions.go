package models

import "time"

type Question struct {
	ID         uint  `gorm:"primaryKey"`
	QuizID     uint  `gorm:"not null"`
	Quiz       *Quiz `gorm:"foreignKey:QuizID;references:ID"`
	Text       string
	Type       string       `gorm:"type:question_type;default:'libre'"` //libre, choix_multiple
	Answer     []Answer     `gorm:"foreignKey:QuestionID;references:ID"`
	UserAnswer []UserAnswer `gorm:"foreignKey:QuestionID;references:ID"`
	CreatedAt  time.Time    `gorm:"autoCreateTime"`
	UpdatedAt  time.Time    `gorm:"autoUpdateTime"`
}

type QuestionForm struct {
	Text    string       `json:"text" binding:"required"`
	Type    string       `json:"type" binding:"required"`
	Answers []AnswerForm `json:"answers" binding:"required"`
}
