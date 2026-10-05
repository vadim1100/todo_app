package tasks_service

import (
	"context"
	"fmt"

	core_domain "github.com/vadim1100/todo_app/internal/core/domain"
	core_redis "github.com/vadim1100/todo_app/internal/core/repository/redis"
)

func (s *TasksService) Create(
	ctx context.Context,
	userID int,
	input CreateTaskInput,
	) (core_domain.Task, error) {
	if err := validateTitle(input.Title); err != nil {
		return core_domain.Task{}, err
	}
	if err := validateDescription(input.Description); err != nil{
		return core_domain.Task{}, err
	}

	task := core_domain.NewTaskUninitialized(userID, input.Title, input.Description)

	createdTask, err := s.tasksRepository.Create(ctx, task)
	if err != nil {
		return core_domain.Task{}, fmt.Errorf("create task in repository: %w", err)
	}

	_ = s.cache.Delete(
		ctx,
		core_redis.NewTasksListKey(createdTask.UserID),
	)

	return createdTask, nil
}