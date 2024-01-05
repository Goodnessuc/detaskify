package scheduler

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ScheduleExecutionLog struct {
	gorm.Model `json:"-"`
	ScheduleID uuid.UUID `json:"schedule_id"`
	ExecutedAt time.Time `json:"executed_at" validate:"required"`
	Status     string    `json:"status" validate:"required,max=50"`
	Output     string    `json:"output"`
}

type ExecutionLogService interface {
	LogExecution(ctx context.Context, taskLog *ScheduleExecutionLog) error
	GetExecutionLog(ctx context.Context, logID uint) (*ScheduleExecutionLog, error)
	UpdateExecutionLog(ctx context.Context, logID uint, updatedLog *ScheduleExecutionLog) error
	DeleteExecutionLog(ctx context.Context, logID uint) error
	ListExecutionLogsByScheduleID(ctx context.Context, scheduleID uuid.UUID) ([]ScheduleExecutionLog, error)
}

// ExecutionLogRepository is the blueprint for execution log-related logic
type ExecutionLogRepository struct {
	service ExecutionLogService
}

// NewExecutionLogService creates a new execution log service
func NewExecutionLogService(service ExecutionLogService) ExecutionLogRepository {
	return ExecutionLogRepository{
		service: service,
	}
}

func (e *ExecutionLogRepository) LogExecution(ctx context.Context, taskLog *ScheduleExecutionLog) error {
	return e.service.LogExecution(ctx, taskLog)
}

func (e *ExecutionLogRepository) GetExecutionLog(ctx context.Context, logID uint) (*ScheduleExecutionLog, error) {
	return e.service.GetExecutionLog(ctx, logID)
}

func (e *ExecutionLogRepository) UpdateExecutionLog(ctx context.Context, logID uint, updatedLog *ScheduleExecutionLog) error {
	return e.service.UpdateExecutionLog(ctx, logID, updatedLog)
}

func (e *ExecutionLogRepository) DeleteExecutionLog(ctx context.Context, logID uint) error {
	return e.service.DeleteExecutionLog(ctx, logID)
}

func (e *ExecutionLogRepository) ListExecutionLogsByScheduleID(ctx context.Context, scheduleID uuid.UUID) ([]ScheduleExecutionLog, error) {
	return e.service.ListExecutionLogsByScheduleID(ctx, scheduleID)
}
