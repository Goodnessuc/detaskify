package db

import (
	"context"
	"gorm.io/gorm"
	"log"
	"time"
)

// Reminder represents a reminder related to a task
type Reminder struct {
	gorm.Model
	Task    Task      `gorm:"foreignKey:ID"`
	Creator string    `gorm:"not null"`
	Message string    `gorm:"size:255;not null"`
	Time    time.Time `gorm:"not null"`
	Repeat  []string  `gorm:"type:text[]"`
	Status  string    `gorm:"size:50;not null"`
}

// CreateReminder creates a new reminder
func (d *Database) CreateReminder(ctx context.Context, reminder *Reminder) error {
	err := d.Client.WithContext(ctx).Create(reminder).Error
	if err != nil {
		log.Printf("Error creating reminder: %s", err.Error())
		return err
	}
	return nil
}

// GetReminder retrieves a reminder by its ID
func (d *Database) GetReminder(ctx context.Context, id uint) (*Reminder, error) {
	var reminder Reminder
	err := d.Client.WithContext(ctx).First(&reminder, id).Error
	if err != nil {
		log.Printf("Error getting reminder: %s", err.Error())
		return nil, err
	}
	return &reminder, nil
}

// UpdateReminder updates an existing reminder
func (d *Database) UpdateReminder(ctx context.Context, id uint, updateData *Reminder) error {
	result := d.Client.WithContext(ctx).Model(&Reminder{}).Where("id = ?", id).Updates(updateData)
	if result.Error != nil {
		log.Printf("Error updating reminder: %s", result.Error.Error())
		return result.Error
	}
	return nil
}

// DeleteReminder deletes a reminder by its ID
func (d *Database) DeleteReminder(ctx context.Context, id uint) error {
	result := d.Client.WithContext(ctx).Where("id = ?", id).Delete(&Reminder{})
	if result.Error != nil {
		log.Printf("Error deleting reminder: %s", result.Error.Error())
		return result.Error
	}
	return nil
}
