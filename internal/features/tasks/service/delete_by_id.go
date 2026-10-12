package tasks_service

import (
	"context"
	"strconv"
	"time"

	core_kafka "github.com/vadim1100/todo_app/internal/core/repository/kafka"
	core_redis "github.com/vadim1100/todo_app/internal/core/repository/redis"
	"go.uber.org/zap"
)

func (s *TasksService) DeleteByTaskID(
	ctx context.Context,
	taskID int,
	userID int,
) error {
	if err := s.tasksRepository.DeleteByTaskID(ctx, taskID, userID); err != nil {
		return err
	}

	if err := s.cache.Delete(
		ctx,
		core_redis.NewTasksListKey(userID),
		core_redis.NewTaskItemKey(taskID),
	); err != nil {
		s.logger.Warn("failed to invalidate task cache",
			zap.Int("user_id", userID),
			zap.Int("task_id", taskID),
			zap.Error(err),
		)
	} else {
		s.logger.Debug("task cache invalidated",
			zap.Int("user_id", userID),
			zap.Int("task_id", taskID),
		)
	}

	if err := s.producer.Produce(ctx, strconv.Itoa(userID), core_kafka.NewTaskEvent(
		core_kafka.EventTaskDeleted,
		userID,
		taskID,
		time.Now().UTC(),
	)); err != nil {
		s.logger.Warn("failed to produce task_deleted event",
			zap.Int("task_id", taskID),
			zap.Error(err),
		)
	} else {
		s.logger.Debug("task_deleted event produced",
			zap.Int("task_id", taskID),
		)
	}

	return nil
}
