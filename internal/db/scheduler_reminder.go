package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ScheduleTaskReminders struct {
	gorm.Model
	ScheduleID   uuid.UUID `gorm:"index;not null"` // Foreign key to Schedule
	ReminderTime time.Time `gorm:"type:datetime;not null"`
	Type         string    `gorm:"type:varchar(50)"`
	Message      string    `gorm:"type:text"`
	Status       string    `gorm:"type:varchar(50)"`
}
