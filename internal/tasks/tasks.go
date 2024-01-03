package tasks

import (
	"github.com/google/uuid"
	"time"
)

type Task struct {
	ID              uuid.UUID `json:"id" validate:"required"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
	Group           string    `json:"group" validate:"max=255"`
	CreatorUsername string    `json:"creator_username" validate:"required,max=255"`
	Title           string    `json:"title" validate:"required,max=255"`
	Description     string    `json:"description" validate:"max=1024"`
	Priority        string    `json:"priority" validate:"max=50"`
	Type            string    `json:"type" validate:"max=50"`
	Deadline        time.Time `json:"deadline"`
	Status          string    `json:"status" validate:"required,max=50"`
	Assignees       []string  `json:"assignees"`
	Tags            []string  `json:"tags"`
}
