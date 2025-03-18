package models

type Quiz struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null;unique"`
	Description string
	TeacherID   uint       `gorm:"column:teacher_id;not null"`
	Teacher     *User      `gorm:"foreignKey:TeacherID;references:ID"`
	Duration    uint       `gorm:"not null"` //minutes
	Questions   []Question `gorm:"foreignKey:QuizID;references:ID"`
	Sessions    []Session  `gorm:"foreignKey:QuizID;references:ID"`
}
