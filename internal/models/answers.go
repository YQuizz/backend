package models

type Answer struct {
	ID         uint      `gorm:";primaryKey"`
	QuestionID uint      `gorm:"not null"`
	Question   *Question `gorm:"foreignKey:QuestionID;references:ID"`
	Text       string    `gorm:"not null"`
	IsCorrect  bool      `gorm:"not null;default:false"`
}
