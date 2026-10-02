package tasks_service

import (
	"context"

	core_domain "github.com/vadim1100/todo_app/internal/core/domain"
)

func (s *TasksService) GetByTaskID (
	ctx context.Context,
	taskID int,
	userID int,
) (core_domain.Task, error){
	return s.tasksRepository.GetByTaskID(ctx, taskID, userID)
}