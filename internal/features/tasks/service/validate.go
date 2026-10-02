package tasks_service

import (
	"fmt"

	core_errors "github.com/vadim1100/todo_app/internal/core/errors"
)

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