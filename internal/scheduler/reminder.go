package scheduler

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ScheduleReminder struct {
	gorm.Model   `json:"-"`
	ScheduleID   uuid.UUID `json:"schedule_id"`
	ReminderTime time.Time `json:"reminder_time" validate:"required"`
	Type         string    `json:"type" validate:"required,max=50"`
	Message      string    `json:"message"`
	Status       string    `json:"status" validate:"required,max=50"`
}

type ReminderService interface {
	CreateScheduleReminder(ctx context.Context, reminder *ScheduleReminder) error
	GetScheduleReminder(ctx context.Context, reminderID uint) (*ScheduleReminder, error)
	UpdateScheduleReminder(ctx context.Context, reminderID uint, reminder *ScheduleReminder) error
	DeleteScheduleReminder(ctx context.Context, reminderID uint) error
	ListRemindersInAscendingOrder(ctx context.Context) ([]ScheduleReminder, error)
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

func (r *ReminderRepository) CreateScheduleReminder(ctx context.Context, reminder *ScheduleReminder) error {
	return r.service.CreateScheduleReminder(ctx, reminder)
}

func (r *ReminderRepository) GetScheduleReminder(ctx context.Context, reminderID uint) (*ScheduleReminder, error) {
	return r.service.GetScheduleReminder(ctx, reminderID)
}

func (r *ReminderRepository) UpdateScheduleReminder(ctx context.Context, reminderID uint, reminder *ScheduleReminder) error {
	return r.service.UpdateScheduleReminder(ctx, reminderID, reminder)
}

func (r *ReminderRepository) DeleteScheduleReminder(ctx context.Context, reminderID uint) error {
	return r.service.DeleteScheduleReminder(ctx, reminderID)
}

func (r *ReminderRepository) ListRemindersInAscendingOrder(ctx context.Context) ([]ScheduleReminder, error) {
	return r.service.ListRemindersInAscendingOrder(ctx)
}
