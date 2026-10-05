package tasks_service

import (
	"context"
	"encoding/json"

	core_domain "github.com/vadim1100/todo_app/internal/core/domain"
	core_redis "github.com/vadim1100/todo_app/internal/core/repository/redis"
)

func (s *TasksService) ListByUser(
	ctx context.Context,
	userID int,
) ([]core_domain.Task, error){
	key := core_redis.NewTasksListKey(userID)
	
	if raw, err := s.cache.Get(ctx, key); err == nil {
		var tasks []core_domain.Task
		if err := json.Unmarshal(raw, &tasks); err == nil {
			return tasks, nil
		}
	}

	tasks, err := s.tasksRepository.ListByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	if raw, err := json.Marshal(tasks); err == nil {
		_ = s.cache.Set(ctx, key, raw, tasksListTTL)
	}

	return tasks, nil
}