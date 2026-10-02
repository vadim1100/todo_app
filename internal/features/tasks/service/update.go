package tasks_service

import (
	"context"
	"fmt"
	"time"

	core_domain "github.com/vadim1100/todo_app/internal/core/domain"
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
			timeNow := time.Now()
			task.CompletedAt = &timeNow
		}
		
		if !*input.Completed && task.Completed{
			task.CompletedAt = nil
		}
		task.Completed = *input.Completed
	}

	updatedTask, err := s.tasksRepository.Update(ctx, task)
	if err != nil {
		return core_domain.Task{}, fmt.Errorf("update task in repository: %w", err)
	}
	
	return updatedTask, nil
}