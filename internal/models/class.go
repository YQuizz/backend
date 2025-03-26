package models

import "time"

type Class struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null;unique"`
	Users     []User    `gorm:"foreignKey:ClassID;references:ID"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
