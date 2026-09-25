package tasks_service

import (
	"context"

	core_domain "github.com/vadim1100/todo_app/internal/core/domain"
)

type TasksService struct {
	tasksRepository TasksRepository
}

type TasksRepository interface{
	Create(
		ctx context.Context,
		task core_domain.Task,
	) (core_domain.Task, error)
}

func NewTasksService(tasksRepository TasksRepository) *TasksService {
	return &TasksService{
		tasksRepository: tasksRepository,
	}
}