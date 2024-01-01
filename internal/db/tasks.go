package db

import (
	"gorm.io/gorm"
	"time"
)

// Task represents a task with various attributes
type Task struct {
	gorm.Model
	Group       string `gorm:"size:255"`
	Title       string `gorm:"size:255;not null"`
	Description string `gorm:"size:1024"`
	Priority    string `gorm:"size:50"`
	Type        string `gorm:"size:50"`
	Deadline    time.Time
	Status      string     `gorm:"size:50;not null"`
	Assignees   []string   `gorm:"type:text[]"`
	Tags        []string   `gorm:"type:text[]"`
	Reminder    []Reminder `gorm:"foreignKey:TaskID"`
	Comments    []Comment  `gorm:"foreignKey:TaskID"`
}

// Reminder represents a reminder related to a task
type Reminder struct {
	gorm.Model
	TaskID  uint      `gorm:"not null"`
	Message string    `gorm:"size:255;not null"`
	Time    time.Time `gorm:"not null"`
	Repeat  []string  `gorm:"type:text[]"`
}

// Comment represents a comment made on a task
type Comment struct {
	gorm.Model
	TaskID   uint   `gorm:"not null"`
	UserID   string `gorm:"size:255;not null"`
	Username string `gorm:"size:255;not null"`
	Text     string `gorm:"size:1024;not null"`
}
