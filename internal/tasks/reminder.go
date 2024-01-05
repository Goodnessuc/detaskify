package tasks

import (
	"context"
	"time"
)

// TaskReminders represents a reminder related to a task
type TaskReminders struct {
	ID      uint      `json:"id"`
	TaskID  uint      `json:"task_id" validate:"required"`
	Creator string    `json:"creator" validate:"required"`
	Message string    `json:"message" validate:"required,max=255"`
	Time    time.Time `json:"time" validate:"required"`
	Repeat  []string  `json:"repeat"`
	Status  string    `json:"status" validate:"required,max=50"`
}

type ReminderService interface {
	CreateReminder(ctx context.Context, reminder *TaskReminders) error
	GetReminder(ctx context.Context, id uint) (*TaskReminders, error)
	UpdateReminder(ctx context.Context, id uint, reminder *TaskReminders) error
	DeleteReminder(ctx context.Context, id uint) error
	ListRemindersByTaskID(ctx context.Context, taskID uint) ([]TaskReminders, error)
}

// ReminderRepository is the blueprint for reminder-related logic
type ReminderRepository struct {
	service ReminderService
}

// NewReminderService creates a new reminder service
func NewReminderService(service ReminderService) ReminderRepository {
	return ReminderRepository{
		service: service,
	}
}

func (r *ReminderRepository) CreateReminder(ctx context.Context, reminder *TaskReminders) error {
	return r.service.CreateReminder(ctx, reminder)
}

func (r *ReminderRepository) GetReminder(ctx context.Context, id uint) (*TaskReminders, error) {
	return r.service.GetReminder(ctx, id)
}

func (r *ReminderRepository) UpdateReminder(ctx context.Context, id uint, reminder *TaskReminders) error {
	return r.service.UpdateReminder(ctx, id, reminder)
}

func (r *ReminderRepository) DeleteReminder(ctx context.Context, id uint) error {
	return r.service.DeleteReminder(ctx, id)
}

func (r *ReminderRepository) ListRemindersByTaskID(ctx context.Context, taskID uint) ([]TaskReminders, error) {
	return r.service.ListRemindersByTaskID(ctx, taskID)
}
