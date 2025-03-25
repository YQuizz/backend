package models

type User struct {
	ID             uint            `json:"id" gorm:"column:id;primaryKey"`
	Email          string          `json:"email" gorm:"unique;not null"`
	Password       string          `json:"-" gorm:"not null"`
	FirstName      string          `json:"first_name" gorm:"not null"`
	LastName       string          `json:"last_name" gorm:"not null"`
	Role           string          `json:"role" gorm:"type:user_role;default:'student'"` //student, teacher, admin
	ClassID        *uint           `json:"class_id"`
	Class          *Class          `gorm:"foreignKey:ClassID;references:ID"`
	Quizzes        []Quiz          `gorm:"foreignKey:TeacherID;references:ID"`
	UserAnswers    []UserAnswer    `gorm:"foreignKey:UserID;references:ID"`
	MonitoringLogs []MonitoringLog `gorm:"foreignKey:UserID;references:ID"`
}

type LoginForm struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}
