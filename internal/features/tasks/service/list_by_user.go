package tasks_service

import (
	"context"
	"encoding/json"

	core_domain "github.com/vadim1100/todo_app/internal/core/domain"
	core_redis "github.com/vadim1100/todo_app/internal/core/repository/redis"
	"go.uber.org/zap"
)

func (s *TasksService) ListByUser(
	ctx context.Context,
	userID int,
) ([]core_domain.Task, error) {
	key := core_redis.NewTasksListKey(userID)

	if raw, err := s.cache.Get(ctx, key); err == nil {
		var tasks []core_domain.Task
		if err := json.Unmarshal(raw, &tasks); err == nil {
			s.logger.Debug("cache hit: tasks list",
				zap.Int("user_id", userID),
				zap.Int("count", len(tasks)),
			)
			return tasks, nil
		}
		s.logger.Warn("cache corrupted, fallback to db",
			zap.Int("user_id", userID),
		)
	} else {
		s.logger.Debug("cache miss: tasks list", zap.Int("user_id", userID))
	}

	tasks, err := s.tasksRepository.ListByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	raw, err := json.Marshal(tasks)
	if err != nil {
		s.logger.Warn("failed to marshal tasks list",
			zap.Int("user_id", userID),
			zap.Error(err),
		)
		return tasks, nil
	}

	if err := s.cache.Set(ctx, key, raw, tasksListTTL); err != nil {
		s.logger.Warn("failed to cache tasks list",
			zap.Int("user_id", userID),
			zap.Error(err),
		)
	}

	return tasks, nil
}
