package tasks_service

import (
	"fmt"
	"strings"

	core_errors "github.com/vadim1100/todo_app/internal/core/errors"
)

func validateTitle(title string) error {
	trimmedTitle := strings.TrimSpace(title)
	titleLen := len([]rune(trimmedTitle))

	if len(trimmedTitle) == 0 {
		return fmt.Errorf(
			"%w: title is empty",
			core_errors.ErrInvalidArgument,
		)
	}
	if titleLen > 100 {
		return fmt.Errorf(
			"%w: title must be <100 characters",
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
