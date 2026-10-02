package tasks_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/vadim1100/todo_app/internal/core/errors"
)

func (r *TasksRepository) DeleteByTaskID(
	ctx context.Context,
	taskID int,
	userID int,
) error {
	query := `
	DELETE FROM todoapp.tasks WHERE id = $1 AND user_id = $2
	`

	commandTag, err := r.pool.Exec(ctx, query, taskID, userID)
	if err != nil {
		return fmt.Errorf("delete task: %w", err)
	}

	if commandTag.RowsAffected() == 0{
		return core_errors.ErrNotFound
	}

	return nil
}