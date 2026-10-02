package tasks_http_transport

import (
	"time"

	core_domain "github.com/vadim1100/todo_app/internal/core/domain"
)

type TaskResponse struct {
	ID          int	`json:"id"`
	Title       string	`json:"title"`
	Description *string	`json:"description,omitempty"`
	Completed   bool	`json:"completed"`
	CreatedAt   time.Time	`json:"created_at"`
	CompletedAt *time.Time	`json:"completed_at,omitempty"`
}

func NewTaskResponse(
	task core_domain.Task,
) TaskResponse {
	return TaskResponse{
		ID: task.ID,
		Title: task.Title,
		Description: task.Description,
		Completed: task.Completed,
		CreatedAt: task.CreatedAt,
		CompletedAt: task.CompletedAt,
	}
}