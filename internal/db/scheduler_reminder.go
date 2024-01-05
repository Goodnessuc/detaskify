package db

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"time"
)

type ScheduleReminders struct {
	gorm.Model
	ScheduleID   uuid.UUID `gorm:"index;not null"` // Foreign key to Schedule
	ReminderTime time.Time `gorm:"type:datetime;not null"`
	Type         string    `gorm:"type:varchar(50)"`
	Message      string    `gorm:"type:text"`
	Status       string    `gorm:"type:varchar(50)"`
}

func (d *Database) CreateScheduleReminder(ctx context.Context, reminder *ScheduleReminders) error {
	err := d.Client.WithContext(ctx).Create(reminder).Error
	if err != nil {
		log.Printf("Error creating reminder: %s", err.Error())
		return err
	}
	return nil
}

func (d *Database) GetScheduleReminder(ctx context.Context, reminderID uint) (*ScheduleReminders, error) {
	var reminder ScheduleReminders
	err := d.Client.WithContext(ctx).Where("id = ?", reminderID).First(&reminder).Error
	if err != nil {
		log.Printf("Error retrieving reminder: %s", err.Error())
		return nil, err
	}
	return &reminder, nil
}

func (d *Database) UpdateScheduleReminder(ctx context.Context, reminderID uint, updatedReminder *ScheduleReminders) error {
	err := d.Client.WithContext(ctx).Model(&ScheduleReminders{}).Where("id = ?", reminderID).Updates(updatedReminder).Error
	if err != nil {
		log.Printf("Error updating reminder: %s", err.Error())
		return err
	}
	return nil
}

func (d *Database) DeleteScheduleReminder(ctx context.Context, reminderID uint) error {
	result := d.Client.WithContext(ctx).Where("id = ?", reminderID).Delete(&ScheduleReminders{})
	if result.Error != nil {
		log.Printf("Error deleting reminder: %s", result.Error.Error())
		return result.Error
	}
	return nil
}

func (d *Database) ListRemindersInAscendingOrder(ctx context.Context) ([]ScheduleReminders, error) {
	var reminders []ScheduleReminders
	err := d.Client.WithContext(ctx).Order("reminder_time asc").Find(&reminders).Error
	if err != nil {
		log.Printf("Error listing reminders in ascending order: %s", err.Error())
		return nil, err
	}
	return reminders, nil
}
