package db

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
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

func (d *Database) CreateSchedule(ctx context.Context, schedule *Schedule) error {
	err := d.Client.WithContext(ctx).Create(schedule).Error
	if err != nil {
		log.Printf("Error creating schedule: %s", err.Error())
		return err
	}
	return nil
}

func (d *Database) GetSchedule(ctx context.Context, scheduleID uuid.UUID) (*Schedule, error) {
	var schedule Schedule
	err := d.Client.WithContext(ctx).Where("id = ?", scheduleID).First(&schedule).Error
	if err != nil {
		log.Printf("Error retrieving schedule: %s", err.Error())
		return nil, err
	}
	return &schedule, nil
}

func (d *Database) UpdateSchedule(ctx context.Context, scheduleID uuid.UUID, updatedSchedule *Schedule) error {
	err := d.Client.WithContext(ctx).Model(&Schedule{}).Where("id = ?", scheduleID).Updates(updatedSchedule).Error
	if err != nil {
		log.Printf("Error updating schedule: %s", err.Error())
		return err
	}
	return nil
}

func (d *Database) DeleteSchedule(ctx context.Context, scheduleID uuid.UUID) error {
	result := d.Client.WithContext(ctx).Where("id = ?", scheduleID).Delete(&Schedule{})
	if result.Error != nil {
		log.Printf("Error deleting schedule: %s", result.Error.Error())
		return result.Error
	}
	return nil
}

func (d *Database) ListSchedules(ctx context.Context) ([]Schedule, error) {
	var schedules []Schedule
	err := d.Client.WithContext(ctx).Find(&schedules).Error
	if err != nil {
		log.Printf("Error listing schedules: %s", err.Error())
		return nil, err
	}
	return schedules, nil
}
