package tasks_service

import (
	"context"

	core_domain "github.com/vadim1100/todo_app/internal/core/domain"
)

func (s *TasksService) ListByUser(
	ctx context.Context,
	userID int,
) ([]core_domain.Task, error){
	return s.tasksRepository.ListByUser(ctx, userID)
}