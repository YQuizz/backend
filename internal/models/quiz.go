package models

import "time"

type Quiz struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null;unique"`
	Description string
	TeacherID   uint       `gorm:"column:teacher_id;not null"`
	Teacher     *User      `gorm:"foreignKey:TeacherID;references:ID"`
	Duration    uint       `gorm:"not null"` //minutes
	Questions   []Question `gorm:"foreignKey:QuizID;references:ID"`
	Sessions    []Session  `gorm:"foreignKey:QuizID;references:ID"`
	CreatedAt   time.Time  `gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime"`
}

type QuizzForm struct {
	Title       string         `json:"title" binding:"required"`
	Description string         `json:"description" binding:"required"`
	Duration    uint           `json:"duration" binding:"required"`
	Questions   []QuestionForm `json:"questions" binding:"required"`
}
