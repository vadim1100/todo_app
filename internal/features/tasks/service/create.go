package tasks_service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	core_domain "github.com/vadim1100/todo_app/internal/core/domain"
	core_kafka "github.com/vadim1100/todo_app/internal/core/repository/kafka"
	core_redis "github.com/vadim1100/todo_app/internal/core/repository/redis"
	"go.uber.org/zap"
)

func (s *TasksService) Create(
	ctx context.Context,
	userID int,
	input CreateTaskInput,
) (core_domain.Task, error) {
	if err := validateTitle(input.Title); err != nil {
		return core_domain.Task{}, err
	}
	if err := validateDescription(input.Description); err != nil {
		return core_domain.Task{}, err
	}

	task := core_domain.NewTaskUninitialized(userID, input.Title, input.Description)

	createdTask, err := s.tasksRepository.Create(ctx, task)
	if err != nil {
		return core_domain.Task{}, fmt.Errorf("create task in repository: %w", err)
	}

	if err := s.cache.Delete(
		ctx,
		core_redis.NewTasksListKey(createdTask.UserID),
	); err != nil {
		s.logger.Warn("failed to invalidate tasks list cache",
			zap.Int("user_id", createdTask.UserID),
			zap.Error(err),
		)
	} else {
		s.logger.Debug("tasks list cache invalidated",
			zap.Int("user_id", createdTask.UserID),
		)
	}

	if err := s.producer.Produce(ctx, strconv.Itoa(userID), core_kafka.NewTaskEvent(
		core_kafka.EventTaskCreated,
		userID,
		createdTask.ID,
		time.Now().UTC(),
	)); err != nil {
		s.logger.Warn("failed to produce task_created event",
			zap.Int("task_id", createdTask.ID),
			zap.Error(err),
		)
	} else {
		s.logger.Debug("task_created event produced",
			zap.Int("task_id", createdTask.ID),
		)
	}

	return createdTask, nil
}
