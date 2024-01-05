package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Schedule struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt          `gorm:"index"`
	UserID      uint                    `gorm:"index;not null"` // Foreign key to User
	Title       string                  `gorm:"type:varchar(100);not null"`
	Description string                  `gorm:"type:text"`
	Priority    string                  `gorm:"type:varchar(50)"`
	Type        string                  `gorm:"type:varchar(50)"`
	Deadline    time.Time               `gorm:"type:datetime"`
	Command     string                  `gorm:"type:text"`
	Status      string                  `gorm:"type:varchar(50)"`
	NextRunTime time.Time               `gorm:"type:datetime"`
	Reminders   []ScheduleTaskReminders `gorm:"foreignKey:ScheduleID"`
}
