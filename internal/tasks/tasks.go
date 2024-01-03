package tasks

import (
	"context"
	"github.com/google/uuid"
	"time"
)

// Task represents a task with various attributes
type Task struct {
	ID              uuid.UUID     `json:"id"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
	Group           string        `json:"group" validate:"max=255"`
	CreatorUsername string        `json:"creator_username" validate:"required,max=255"`
	Title           string        `json:"title" validate:"required,max=255"`
	Description     string        `json:"description" validate:"max=1024"`
	Priority        string        `json:"priority" validate:"max=50"`
	Type            string        `json:"type" validate:"max=50"`
	Deadline        time.Time     `json:"deadline"`
	Status          string        `json:"status" validate:"required,max=50"`
	Assignees       []string      `json:"assignees"`
	Tags            []string      `json:"tags"`
	Comments        []TaskComment `json:"-"`
	Reminders       []Reminder    `json:"-"`
}

type TaskService interface {
	CreateTask(ctx context.Context, task *Task) error
	GetTask(ctx context.Context, id uint) (*Task, error)
	UpdateTask(ctx context.Context, id uint, task *Task) error
	DeleteTask(ctx context.Context, id uint) error
	GetUserTasks(ctx context.Context, username string) ([]Task, error)
	SearchTasks(ctx context.Context, title, description, priority string) ([]Task, error)
	AddAssigneeToTask(ctx context.Context, taskID uint, assignee string) error
	RemoveAssigneeFromTask(ctx context.Context, taskID uint, assignee string) error
	ListTasksByStatus(ctx context.Context, status string) ([]Task, error)
	ListTasksByDeadline(ctx context.Context, start, end time.Time) ([]Task, error)
	ListOverdueTasks(ctx context.Context, currentDate time.Time) ([]Task, error)
	AddTagToTask(ctx context.Context, taskID uint, tag string) error
	RemoveTagFromTask(ctx context.Context, taskID uint, tag string) error
	ListTasksByPriority(ctx context.Context, priority string) ([]Task, error)
	ListTasksForReminder(ctx context.Context, start, end time.Time) ([]Task, error)
}

// TaskRepository is the blueprint for task-related logic
type TaskRepository struct {
	service TaskService
}

// NewTaskService creates a new task service
func NewTaskService(service TaskService) TaskRepository {
	return TaskRepository{
		service: service,
	}
}

func (t *TaskRepository) CreateTask(ctx context.Context, task *Task) error {
	return t.service.CreateTask(ctx, task)
}

func (t *TaskRepository) GetTask(ctx context.Context, id uint) (*Task, error) {
	return t.service.GetTask(ctx, id)
}

func (t *TaskRepository) UpdateTask(ctx context.Context, id uint, task *Task) error {
	return t.service.UpdateTask(ctx, id, task)
}

func (t *TaskRepository) DeleteTask(ctx context.Context, id uint) error {
	return t.service.DeleteTask(ctx, id)
}

func (t *TaskRepository) GetUserTasks(ctx context.Context, username string) ([]Task, error) {
	return t.service.GetUserTasks(ctx, username)
}

func (t *TaskRepository) SearchTasks(ctx context.Context, title, description, priority string) ([]Task, error) {
	return t.service.SearchTasks(ctx, title, description, priority)
}

func (t *TaskRepository) AddAssigneeToTask(ctx context.Context, taskID uint, assignee string) error {
	return t.service.AddAssigneeToTask(ctx, taskID, assignee)
}

func (t *TaskRepository) RemoveAssigneeFromTask(ctx context.Context, taskID uint, assignee string) error {
	return t.service.RemoveAssigneeFromTask(ctx, taskID, assignee)
}

func (t *TaskRepository) ListTasksByStatus(ctx context.Context, status string) ([]Task, error) {
	return t.service.ListTasksByStatus(ctx, status)
}

func (t *TaskRepository) ListTasksByDeadline(ctx context.Context, start, end time.Time) ([]Task, error) {
	return t.service.ListTasksByDeadline(ctx, start, end)
}

func (t *TaskRepository) ListOverdueTasks(ctx context.Context, currentDate time.Time) ([]Task, error) {
	return t.service.ListOverdueTasks(ctx, currentDate)
}

func (t *TaskRepository) AddTagToTask(ctx context.Context, taskID uint, tag string) error {
	return t.service.AddTagToTask(ctx, taskID, tag)
}

func (t *TaskRepository) RemoveTagFromTask(ctx context.Context, taskID uint, tag string) error {
	return t.service.RemoveTagFromTask(ctx, taskID, tag)
}

func (t *TaskRepository) ListTasksByPriority(ctx context.Context, priority string) ([]Task, error) {
	return t.service.ListTasksByPriority(ctx, priority)
}

func (t *TaskRepository) ListTasksForReminder(ctx context.Context, start, end time.Time) ([]Task, error) {
	return t.service.ListTasksForReminder(ctx, start, end)
}
