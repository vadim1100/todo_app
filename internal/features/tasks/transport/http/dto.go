package tasks_http_transport

import (
	"time"

	core_domain "github.com/vadim1100/todo_app/internal/core/domain"
)

type TaskResponse struct {
	ID          int	`json:"id" example:"42"`
	Title       string	`json:"title" example:"Сделать домашнее задание"`
	Description *string	`json:"description,omitempty"`
	Completed   bool	`json:"completed" example:"false"`
	CreatedAt   time.Time	`json:"created_at" example:"2026-10-12T15:00:00Z"`
	CompletedAt *time.Time	`json:"completed_at,omitempty"`
}

type ErrorResponse struct {
	Error   string `json:"error" example:"unauthorized"`
	Message string `json:"message" example:"failed to login"`
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