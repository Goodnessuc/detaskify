package db

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"time"
)

// Task represents a task with various attributes
type Task struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	Group           string         `gorm:"size:255"`
	CreatorUsername string         `gorm:"size:255;not null"`
	Title           string         `gorm:"size:255;not null"`
	Description     string         `gorm:"size:1024"`
	Priority        string         `gorm:"size:50"`
	Type            string         `gorm:"size:50"`
	Deadline        time.Time
	Status          string   `gorm:"size:50;not null"`
	Assignees       []string `gorm:"type:text[]"`
	Tags            []string `gorm:"type:text[]"`
	Comments        []TaskComment
	Reminders       []Reminder
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

// GetUserTasks retrieves tasks created by a given user
func (d *Database) GetUserTasks(ctx context.Context, username string) ([]Task, error) {
	var tasks []Task
	err := d.Client.WithContext(ctx).Where("creator_username = ?", username).Find(&tasks).Error
	if err != nil {
		log.Printf("Error getting tasks for user %s: %s", username, err.Error())
		return nil, err
	}
	return tasks, nil
}
