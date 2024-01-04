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

// SearchTasks searches for tasks based on various criteria.
func (d *Database) SearchTasks(ctx context.Context, title, description, priority string) ([]Task, error) {
	var tasks []Task
	err := d.Client.WithContext(ctx).Where("title LIKE ? AND description LIKE ? AND priority = ?", "%"+title+"%", "%"+description+"%", priority).Find(&tasks).Error
	if err != nil {
		log.Printf("Error searching tasks: %s", err.Error())
		return nil, err
	}
	return tasks, nil
}

// AddAssigneeToTask adds an assignee to a task.
func (d *Database) AddAssigneeToTask(ctx context.Context, taskID uint, assignee string) error {
	var task Task
	err := d.Client.WithContext(ctx).First(&task, taskID).Error
	if err != nil {
		log.Printf("Error finding task %d: %s", taskID, err.Error())
		return err
	}
	task.Assignees = append(task.Assignees, assignee)
	return d.Client.WithContext(ctx).Save(&task).Error
}

// RemoveAssigneeFromTask removes an assignee from a task.
func (d *Database) RemoveAssigneeFromTask(ctx context.Context, taskID uint, assignee string) error {
	var task Task
	err := d.Client.WithContext(ctx).First(&task, taskID).Error
	if err != nil {
		log.Printf("Error finding task %d: %s", taskID, err.Error())
		return err
	}

	// Filter out the assignee
	var updatedAssignees []string
	for _, a := range task.Assignees {
		if a != assignee {
			updatedAssignees = append(updatedAssignees, a)
		}
	}
	task.Assignees = updatedAssignees

	return d.Client.WithContext(ctx).Save(&task).Error
}

// ListTasksByStatus lists all tasks with a specific status.
func (d *Database) ListTasksByStatus(ctx context.Context, status string) ([]Task, error) {
	var tasks []Task
	err := d.Client.WithContext(ctx).Where("status = ?", status).Find(&tasks).Error
	if err != nil {
		log.Printf("Error listing tasks with status %s: %s", status, err.Error())
		return nil, err
	}
	return tasks, nil
}

// ListTasksByDeadline lists tasks with deadlines within a certain range.
func (d *Database) ListTasksByDeadline(ctx context.Context, start, end time.Time) ([]Task, error) {
	var tasks []Task
	err := d.Client.WithContext(ctx).Where("deadline BETWEEN ? AND ?", start, end).Find(&tasks).Error
	if err != nil {
		log.Printf("Error listing tasks with deadlines between %s and %s: %s", start, end, err.Error())
		return nil, err
	}
	return tasks, nil
}

// ListOverdueTasks lists tasks that have passed their deadline.
func (d *Database) ListOverdueTasks(ctx context.Context, currentDate time.Time) ([]Task, error) {
	var tasks []Task
	err := d.Client.WithContext(ctx).Where("deadline < ?", currentDate).Find(&tasks).Error
	if err != nil {
		log.Printf("Error listing overdue tasks: %s", err.Error())
		return nil, err
	}
	return tasks, nil
}

// AddTagToTask adds a new tag to a task.
func (d *Database) AddTagToTask(ctx context.Context, taskID uint, tag string) error {
	var task Task
	err := d.Client.WithContext(ctx).First(&task, taskID).Error
	if err != nil {
		log.Printf("Error finding task %d: %s", taskID, err.Error())
		return err
	}

	// Check if the tag already exists
	for _, t := range task.Tags {
		if t == tag {
			return nil // Tag already exists
		}
	}

	task.Tags = append(task.Tags, tag)
	return d.Client.WithContext(ctx).Save(&task).Error
}

// RemoveTagFromTask removes a tag from a task.
func (d *Database) RemoveTagFromTask(ctx context.Context, taskID uint, tag string) error {
	var task Task
	err := d.Client.WithContext(ctx).First(&task, taskID).Error
	if err != nil {
		log.Printf("Error finding task %d: %s", taskID, err.Error())
		return err
	}

	var updatedTags []string
	for _, t := range task.Tags {
		if t != tag {
			updatedTags = append(updatedTags, t)
		}
	}
	task.Tags = updatedTags

	return d.Client.WithContext(ctx).Save(&task).Error
}

// ListTasksByPriority lists tasks by their priority levels.
func (d *Database) ListTasksByPriority(ctx context.Context, priority string) ([]Task, error) {
	var tasks []Task
	err := d.Client.WithContext(ctx).Where("priority = ?", priority).Find(&tasks).Error
	if err != nil {
		log.Printf("Error listing tasks with priority %s: %s", priority, err.Error())
		return nil, err
	}
	return tasks, nil
}

// ListTasksForReminder lists tasks that have reminders set for a specific time range.
func (d *Database) ListTasksForReminder(ctx context.Context, start, end time.Time) ([]Task, error) {
	var tasks []Task
	err := d.Client.WithContext(ctx).Joins("JOIN reminders ON reminders.task_id = tasks.id").
		Where("reminders.time BETWEEN ? AND ?", start, end).Group("tasks.id").Find(&tasks).Error
	if err != nil {
		log.Printf("Error listing tasks with reminders between %s and %s: %s", start, end, err.Error())
		return nil, err
	}
	return tasks, nil
}
