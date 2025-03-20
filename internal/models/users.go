package models

import "yquiz_back/internal/pkg"

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

type UserForm struct {
	Email     string `form:"email" binding:"required,email"`
	Password  string `form:"password" binding:"required,min=8"`
	FirstName string `form:"first_name" binding:"required"`
	LastName  string `form:"last_name" binding:"required"`
	Role      string `form:"role" binding:"required,oneof=student teacher admin"`
}

func (uf *UserForm) ToUser() (*User, error) {
	hashedPassword, err := pkg.HashPassword(uf.Password)
	if err != nil {
		return nil, err
	}

	return &User{
		Email:     uf.Email,
		Password:  hashedPassword,
		FirstName: uf.FirstName,
		LastName:  uf.LastName,
		Role:      uf.Role,
	}, nil
}
