package tasks_service

import (
	"context"
	"encoding/json"

	core_domain "github.com/vadim1100/todo_app/internal/core/domain"
	core_redis "github.com/vadim1100/todo_app/internal/core/repository/redis"
)

func (s *TasksService) GetByTaskID (
	ctx context.Context,
	taskID int,
	userID int,
) (core_domain.Task, error){
	key := core_redis.NewTaskItemKey(taskID)
	
	if raw, err := s.cache.Get(ctx, key); err == nil {
		var task core_domain.Task
		if err := json.Unmarshal(raw, &task); err == nil && task.UserID == userID{
			return task, nil
		}
	}

	task, err := s.tasksRepository.GetByTaskID(ctx, taskID, userID)
	if err != nil {
		return core_domain.Task{}, err
	}
	
	if raw, err := json.Marshal(task); err == nil {
		_ = s.cache.Set(
			ctx,
			key,
			raw,
			taskItemTTL,
		)
	}

	return task, nil
}