package tasks_service

import (
	"context"

	core_redis "github.com/vadim1100/todo_app/internal/core/repository/redis"
)

func (s *TasksService) DeleteByTaskID(
	ctx context.Context,
	taskID int,
	userID int,
) error {
	if err := s.tasksRepository.DeleteByTaskID(ctx, taskID, userID); err != nil {
		return err
	}

	_ = s.cache.Delete(
		ctx,
		core_redis.NewTasksListKey(userID),
		core_redis.NewTaskItemKey(taskID),
	)

	return nil
}