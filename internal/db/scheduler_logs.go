package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ScheduleTaskExecutionLog struct {
	gorm.Model
	ScheduleID uuid.UUID `gorm:"index;not null"`
	ExecutedAt time.Time `gorm:"type:datetime;not null"`
	Status     string    `gorm:"type:varchar(50)"`
	Output     string    `gorm:"type:text"`
}
