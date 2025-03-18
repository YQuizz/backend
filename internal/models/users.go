package models

type User struct {
	ID             uint   `gorm:"column:id;primaryKey"`
	Email          string `gorm:"unique;not null"`
	Password       string `gorm:"not null"`
	FirstName      string `gorm:"not null"`
	LastName       string `gorm:"not null"`
	Role           string `gorm:"type:user_role;default:'student'"` //student, teacher, admin
	ClassID        *uint
	Class          *Class          `gorm:"foreignKey:ClassID;references:ID"`
	Quizzes        []Quiz          `gorm:"foreignKey:TeacherID;references:ID"`
	UserAnswers    []UserAnswer    `gorm:"foreignKey:UserID;references:ID"`
	MonitoringLogs []MonitoringLog `gorm:"foreignKey:UserID;references:ID"`
}
