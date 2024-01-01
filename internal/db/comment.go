package db

import (
	"context"
	"gorm.io/gorm"
	"log"
)

// Comment represents a comment made on a task
type Comment struct {
	gorm.Model
	TaskID   uint   `gorm:"not null"`
	UserID   string `gorm:"size:255;not null"`
	Username string `gorm:"size:255;not null"`
	Text     string `gorm:"size:1024;not null"`
}

// CreateComment creates a new comment
func (d *Database) CreateComment(ctx context.Context, comment *Comment) error {
	err := d.Client.WithContext(ctx).Create(comment).Error
	if err != nil {
		log.Printf("Error creating comment: %s", err.Error())
		return err
	}
	return nil
}

// GetComment retrieves a comment by its ID
func (d *Database) GetComment(ctx context.Context, id uint) (*Comment, error) {
	var comment Comment
	err := d.Client.WithContext(ctx).First(&comment, id).Error
	if err != nil {
		log.Printf("Error getting comment: %s", err.Error())
		return nil, err
	}
	return &comment, nil
}

// UpdateComment updates an existing comment
func (d *Database) UpdateComment(ctx context.Context, id uint, updateData *Comment) error {
	result := d.Client.WithContext(ctx).Model(&Comment{}).Where("id = ?", id).Updates(updateData)
	if result.Error != nil {
		log.Printf("Error updating comment: %s", result.Error.Error())
		return result.Error
	}
	return nil
}

// DeleteComment deletes a comment by its ID
func (d *Database) DeleteComment(ctx context.Context, id uint) error {
	result := d.Client.WithContext(ctx).Where("id = ?", id).Delete(&Comment{})
	if result.Error != nil {
		log.Printf("Error deleting comment: %s", result.Error.Error())
		return result.Error
	}
	return nil
}
