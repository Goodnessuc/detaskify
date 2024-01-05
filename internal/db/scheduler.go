package db

import (
	"gorm.io/gorm"
	"time"
)

// Schedule represents a schedule.
type Schedule struct {
	gorm.Model
	Title       string          `gorm:"type:varchar(100);not null"`
	Description string          `gorm:"type:text"`
	Priority    string          `gorm:"type:varchar(50)"`
	Type        string          `gorm:"type:varchar(50)"`
	Deadline    time.Time       `gorm:"type:datetime"`
	Command     string          `gorm:"type:text"`
	Status      string          `gorm:"type:varchar(50)"`
	Reminders   []TaskReminders `gorm:"foreignKey:ScheduleID"`
}
