package tasks_service

import (
	"context"
	"fmt"

	core_domain "github.com/vadim1100/todo_app/internal/core/domain"
	core_errors "github.com/vadim1100/todo_app/internal/core/errors"
)

func (s *TasksService) Create(ctx context.Context, userID int, input CreateTaskInput) (core_domain.Task, error) {
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
	return createdTask, nil
}

func validateTitle(title string) error {
	titleLen := len([]rune(title))
	if titleLen < 1 || titleLen > 100 {
		return fmt.Errorf(
			"%w: title must be 1-100 characters",
			core_errors.ErrInvalidArgument,
		)
	}
	return nil
}

func validateDescription(description *string) error {
	if description == nil {
		return nil
	}

	descriptionLen := len([]rune(*description))
	if descriptionLen < 1 || descriptionLen > 500 {
		return fmt.Errorf(
			"%w: description must be 1-500 characters",
			core_errors.ErrInvalidArgument,
		)
	}
	return nil
}