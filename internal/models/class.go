package models

type Class struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"not null;unique"`
	Users []User `gorm:"foreignKey:ClassID;references:ID"`
}
