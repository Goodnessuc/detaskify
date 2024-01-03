package tasks

import (
	"context"
	"gorm.io/gorm"
)

// TaskComment represents a comment made on a task
type TaskComment struct {
	gorm.Model
	Task     Task   `json:"task_id" validate:"required"`
	Username string `json:"username" validate:"required,max=255"`
	Text     string `json:"text" validate:"required,max=1024"`
	Status   string `json:"status" validate:"required,max=50"`
	Slug     string `json:"slug" validate:"max=255"`
}

type TaskCommentService interface {
	CreateComment(ctx context.Context, comment *TaskComment) error
	GetComment(ctx context.Context, id uint) (*TaskComment, error)
	UpdateComment(ctx context.Context, id uint, comment *TaskComment) error
	DeleteComment(ctx context.Context, id uint) error
	ListCommentsByTaskID(ctx context.Context, taskID uint) ([]TaskComment, error)
}

// TaskCommentRepository is the blueprint for comment-related logic
type TaskCommentRepository struct {
	service TaskCommentService
}

// NewTaskCommentService creates a new task comment service
func NewTaskCommentService(service TaskCommentService) TaskCommentRepository {
	return TaskCommentRepository{
		service: service,
	}
}

func (t *TaskCommentRepository) CreateComment(ctx context.Context, comment *TaskComment) error {
	return t.service.CreateComment(ctx, comment)
}

func (t *TaskCommentRepository) GetComment(ctx context.Context, id uint) (*TaskComment, error) {
	return t.service.GetComment(ctx, id)
}

func (t *TaskCommentRepository) UpdateComment(ctx context.Context, id uint, comment *TaskComment) error {
	return t.service.UpdateComment(ctx, id, comment)
}

func (t *TaskCommentRepository) DeleteComment(ctx context.Context, id uint) error {
	return t.service.DeleteComment(ctx, id)
}

func (t *TaskCommentRepository) ListCommentsByTaskID(ctx context.Context, taskID uint) ([]TaskComment, error) {
	return t.service.ListCommentsByTaskID(ctx, taskID)
}
