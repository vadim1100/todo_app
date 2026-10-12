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

func (s *TasksService) Update(
	ctx context.Context,
	taskID int,
	userID int,
	input UpdateTaskInput,
) (core_domain.Task, error) {
	task, err := s.tasksRepository.GetByTaskID(ctx, taskID, userID)
	if err != nil {
		return core_domain.Task{}, err
	}

	if input.Title != nil {
		if err := validateTitle(*input.Title); err != nil {
			return core_domain.Task{}, err
		}
		task.Title = *input.Title
	}

	if input.Description != nil {
		if err := validateDescription(input.Description); err != nil {
			return core_domain.Task{}, err
		}
		task.Description = input.Description
	}

	if input.Completed != nil {
		if *input.Completed && !task.Completed {
			timeNow := time.Now().UTC()
			task.CompletedAt = &timeNow
		}

		if !*input.Completed && task.Completed {
			task.CompletedAt = nil
		}
		task.Completed = *input.Completed
	}

	updatedTask, err := s.tasksRepository.Update(ctx, task)
	if err != nil {
		return core_domain.Task{}, fmt.Errorf("update task in repository: %w", err)
	}

	if err := s.cache.Delete(
		ctx,
		core_redis.NewTasksListKey(updatedTask.UserID),
		core_redis.NewTaskItemKey(updatedTask.ID),
	); err != nil {
		s.logger.Warn("failed to invalidate task cache",
			zap.Int("user_id", updatedTask.UserID),
			zap.Int("task_id", updatedTask.ID),
			zap.Error(err),
		)
	} else {
		s.logger.Debug("task cache invalidated",
			zap.Int("user_id", updatedTask.UserID),
			zap.Int("task_id", updatedTask.ID),
		)
	}

	if err := s.producer.Produce(ctx, strconv.Itoa(updatedTask.UserID), core_kafka.NewTaskEvent(
		core_kafka.EventTaskUpdated,
		updatedTask.UserID,
		updatedTask.ID,
		time.Now().UTC(),
	)); err != nil {
		s.logger.Warn("failed to produce task_updated event",
			zap.Int("task_id", updatedTask.ID),
			zap.Error(err),
		)
	} else {
		s.logger.Debug("task_updated event produced",
			zap.Int("task_id", updatedTask.ID),
		)
	}

	return updatedTask, nil
}
