package db

import (
	"context"
	"gorm.io/gorm"
	"log"
	"time"
)

// TaskReminders represents a reminder related to a task
type TaskReminders struct {
	gorm.Model
	TaskID  uint      `gorm:"not null"`
	Task    Task      `gorm:"foreignKey:TaskID"`
	Message string    `gorm:"size:255;not null"`
	Time    time.Time `gorm:"not null"`
	Repeat  []string  `gorm:"type:text[]"`
	Status  string    `gorm:"size:50;not null"`
}

// CreateReminder creates a new reminder
func (d *Database) CreateReminder(ctx context.Context, reminder *TaskReminders) error {
	err := d.Client.WithContext(ctx).Create(reminder).Error
	if err != nil {
		log.Printf("Error creating reminder: %s", err.Error())
		return err
	}
	return nil
}

// GetReminder retrieves a reminder by its ID
func (d *Database) GetReminder(ctx context.Context, id uint) (*TaskReminders, error) {
	var reminder TaskReminders
	err := d.Client.WithContext(ctx).First(&reminder, id).Error
	if err != nil {
		log.Printf("Error getting reminder: %s", err.Error())
		return nil, err
	}
	return &reminder, nil
}

// UpdateReminder updates an existing reminder
func (d *Database) UpdateReminder(ctx context.Context, id uint, updateData *TaskReminders) error {
	result := d.Client.WithContext(ctx).Model(&TaskReminders{}).Where("id = ?", id).Updates(updateData)
	if result.Error != nil {
		log.Printf("Error updating reminder: %s", result.Error.Error())
		return result.Error
	}
	return nil
}

// DeleteReminder deletes a reminder by its ID
func (d *Database) DeleteReminder(ctx context.Context, id uint) error {
	result := d.Client.WithContext(ctx).Where("id = ?", id).Delete(&TaskReminders{})
	if result.Error != nil {
		log.Printf("Error deleting reminder: %s", result.Error.Error())
		return result.Error
	}
	return nil
}

// ListRemindersByTaskID retrieves all reminders for a specific task.
func (d *Database) ListRemindersByTaskID(ctx context.Context, taskID uint) ([]TaskReminders, error) {
	var reminders []TaskReminders
	err := d.Client.WithContext(ctx).Where("task_id = ?", taskID).Find(&reminders).Error
	if err != nil {
		log.Printf("Error retrieving reminders for task %d: %s", taskID, err.Error())
		return nil, err
	}
	return reminders, nil
}
