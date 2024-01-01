package db

import (
	"context"
	"gorm.io/gorm"
	"log"
	"time"
)

// Task represents a task with various attributes
type Task struct {
	gorm.Model
	Group           string `gorm:"size:255"`
	CreatorUsername string `gorm:"size:255;not null"`
	Title           string `gorm:"size:255;not null"`
	Description     string `gorm:"size:1024"`
	Priority        string `gorm:"size:50"`
	Type            string `gorm:"size:50"`
	Deadline        time.Time
	Status          string     `gorm:"size:50;not null"`
	Assignees       []string   `gorm:"type:text[]"`
	Tags            []string   `gorm:"type:text[]"`
	Reminder        []Reminder `gorm:"foreignKey:TaskID"`
	Comments        []Comment  `gorm:"foreignKey:TaskID"`
}

// CreateTask creates a new task in the database
func (d *Database) CreateTask(ctx context.Context, task *Task) error {
	err := d.Client.WithContext(ctx).Create(task).Error
	if err != nil {
		log.Printf("Error creating task: %s", err.Error())
		return err
	}
	return nil
}

// GetTask retrieves a task by its ID
func (d *Database) GetTask(ctx context.Context, id uint) (*Task, error) {
	var task Task
	err := d.Client.WithContext(ctx).Preload("Reminder").Preload("Comments").First(&task, id).Error
	if err != nil {
		log.Printf("Error getting task: %s", err.Error())
		return nil, err
	}
	return &task, nil
}

// UpdateTask updates an existing task
func (d *Database) UpdateTask(ctx context.Context, id uint, updateData *Task) error {
	result := d.Client.WithContext(ctx).Model(&Task{}).Where("id = ?", id).Updates(updateData)
	if result.Error != nil {
		log.Printf("Error updating task: %s", result.Error.Error())
		return result.Error
	}
	return nil
}

// DeleteTask deletes a task by its ID
func (d *Database) DeleteTask(ctx context.Context, id uint) error {
	result := d.Client.WithContext(ctx).Where("id = ?", id).Delete(&Task{})
	if result.Error != nil {
		log.Printf("Error deleting task: %s", result.Error.Error())
		return result.Error
	}
	return nil
}
