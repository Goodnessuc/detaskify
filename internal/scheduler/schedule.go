package scheduler

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type Schedule struct {
	ID          uuid.UUID          `json:"id"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	DeletedAt   *time.Time         `json:"deleted_at"`
	UserID      uint               `json:"user_id"`
	Title       string             `json:"title" validate:"required,max=100"`
	Description string             `json:"description"`
	Priority    string             `json:"priority"`
	Type        string             `json:"type"`
	Deadline    time.Time          `json:"deadline"`
	Command     string             `json:"command"`
	Status      string             `json:"status"`
	NextRunTime time.Time          `json:"next_run_time"`
	Reminders   []ScheduleReminder `json:"reminders"`
}

type ScheduleService interface {
	CreateSchedule(ctx context.Context, schedule *Schedule) error
	GetSchedule(ctx context.Context, scheduleID uuid.UUID) (*Schedule, error)
	UpdateSchedule(ctx context.Context, scheduleID uuid.UUID, schedule *Schedule) error
	DeleteSchedule(ctx context.Context, scheduleID uuid.UUID) error
	GetUserSchedules(ctx context.Context, userID uint) ([]Schedule, error)
	ListSchedules(ctx context.Context) ([]Schedule, error)
}

// ScheduleRepository is the blueprint for schedule-related logic
type ScheduleRepository struct {
	service ScheduleService
}

// NewScheduleService creates a new schedule service
func NewScheduleService(service ScheduleService) ScheduleRepository {
	return ScheduleRepository{
		service: service,
	}
}

func (s *ScheduleRepository) CreateSchedule(ctx context.Context, schedule *Schedule) error {
	return s.service.CreateSchedule(ctx, schedule)
}

func (s *ScheduleRepository) GetSchedule(ctx context.Context, scheduleID uuid.UUID) (*Schedule, error) {
	return s.service.GetSchedule(ctx, scheduleID)
}

func (s *ScheduleRepository) UpdateSchedule(ctx context.Context, scheduleID uuid.UUID, schedule *Schedule) error {
	return s.service.UpdateSchedule(ctx, scheduleID, schedule)
}

func (s *ScheduleRepository) DeleteSchedule(ctx context.Context, scheduleID uuid.UUID) error {
	return s.service.DeleteSchedule(ctx, scheduleID)
}

func (s *ScheduleRepository) GetUserSchedules(ctx context.Context, userID uint) ([]Schedule, error) {
	return s.service.GetUserSchedules(ctx, userID)
}

func (s *ScheduleRepository) ListSchedules(ctx context.Context) ([]Schedule, error) {
	return s.service.ListSchedules(ctx)
}
