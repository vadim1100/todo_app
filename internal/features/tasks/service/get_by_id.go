package tasks_service

import (
	"context"
	"encoding/json"

	core_domain "github.com/vadim1100/todo_app/internal/core/domain"
	core_redis "github.com/vadim1100/todo_app/internal/core/repository/redis"
	"go.uber.org/zap"
)

func (s *TasksService) GetByTaskID(
	ctx context.Context,
	taskID int,
	userID int,
) (core_domain.Task, error) {
	key := core_redis.NewTaskItemKey(taskID)

	if raw, err := s.cache.Get(ctx, key); err == nil {
		var task core_domain.Task
		if err := json.Unmarshal(raw, &task); err == nil && task.UserID == userID {
			s.logger.Debug("cache hit: task",
				zap.Int("task_id", taskID),
				zap.Int("user_id", userID),
			)
			return task, nil
		}
		s.logger.Warn("cache hit but wrong user or corrupted",
			zap.Int("task_id", taskID),
			zap.Int("user_id", userID),
		)
	} else {
		s.logger.Debug("cache miss: task", zap.Int("task_id", taskID))
	}

	task, err := s.tasksRepository.GetByTaskID(ctx, taskID, userID)
	if err != nil {
		return core_domain.Task{}, err
	}

	raw, err := json.Marshal(task)
	if err != nil {
		s.logger.Warn("failed to marshal task",
			zap.Int("task_id", taskID),
			zap.Error(err),
		)
		return task, nil
	}

	if err := s.cache.Set(
		ctx,
		key,
		raw,
		taskItemTTL,
	); err != nil {
		s.logger.Warn("failed to cache task",
			zap.Int("task_id", taskID),
			zap.Error(err),
		)
	}

	return task, nil
}
