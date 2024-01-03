package db

import (
	"context"
	"gorm.io/gorm"
	"log"
)

// TaskComment represents a comment made on a task
type TaskComment struct {
	gorm.Model
	Task     Task   `gorm:"foreignKey:ID"`
	Username string `gorm:"size:255;not null"`
	Text     string `gorm:"size:1024;not null"`
	Status   string `gorm:"size:50;not null"`
	Slug     string `gorm:"size:255"`
}

// CreateComment creates a new comment
func (d *Database) CreateComment(ctx context.Context, comment *TaskComment) error {
	err := d.Client.WithContext(ctx).Create(comment).Error
	if err != nil {
		log.Printf("Error creating comment: %s", err.Error())
		return err
	}
	return nil
}

// GetComment retrieves a comment by its ID
func (d *Database) GetComment(ctx context.Context, id uint) (*TaskComment, error) {
	var comment TaskComment
	err := d.Client.WithContext(ctx).First(&comment, id).Error
	if err != nil {
		log.Printf("Error getting comment: %s", err.Error())
		return nil, err
	}
	return &comment, nil
}

// UpdateComment updates an existing comment
func (d *Database) UpdateComment(ctx context.Context, id uint, updateData *TaskComment) error {
	result := d.Client.WithContext(ctx).Model(&TaskComment{}).Where("id = ?", id).Updates(updateData)
	if result.Error != nil {
		log.Printf("Error updating comment: %s", result.Error.Error())
		return result.Error
	}
	return nil
}

// DeleteComment deletes a comment by its ID
func (d *Database) DeleteComment(ctx context.Context, id uint) error {
	result := d.Client.WithContext(ctx).Where("id = ?", id).Delete(&TaskComment{})
	if result.Error != nil {
		log.Printf("Error deleting comment: %s", result.Error.Error())
		return result.Error
	}
	return nil
}

// ListCommentsByTaskID retrieves all comments for a specific task.
func (d *Database) ListCommentsByTaskID(ctx context.Context, taskID uint) ([]TaskComment, error) {
	var comments []TaskComment
	err := d.Client.WithContext(ctx).Where("task_id = ?", taskID).Find(&comments).Error
	if err != nil {
		log.Printf("Error retrieving comments for task %d: %s", taskID, err.Error())
		return nil, err
	}
	return comments, nil
}
