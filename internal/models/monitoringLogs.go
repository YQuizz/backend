package models

type MonitoringLog struct {
	ID        uint     `gorm:"primaryKey"`
	SessionID uint     `gorm:"not null"`
	UserID    uint     `gorm:"not null"`
	EventType string   `gorm:"type:monitoring_event_type"` // copy, paste, change_tab
	Session   *Session `gorm:"foreignKey:SessionID;references:ID"`
	User      *User    `gorm:"foreignKey:UserID;references:ID"`
}
