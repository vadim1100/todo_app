package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	core_domain "github.com/vadim1100/todo_app/internal/core/domain"
	core_errors "github.com/vadim1100/todo_app/internal/core/errors"
)

func (r *TasksRepository) Update(ctx context.Context, task core_domain.Task) (core_domain.Task, error) {
	query := `
	UPDATE todoapp.tasks
	SET title = $1,
		description = $2,
		completed = $3,
		completed_at = $4,
		version = version + 1
	WHERE id = $5 AND user_id = $6 AND version = $7
	RETURNING version
	`

	row := r.pool.QueryRow(ctx, query,
		task.Title,
		task.Description,
		task.Completed,
		task.CompletedAt,
		task.ID,
		task.UserID,
		task.Version,
	)

	if err := row.Scan(&task.Version); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return core_domain.Task{}, core_errors.ErrConflict
		}
		return core_domain.Task{}, fmt.Errorf("update task: %w", err)
	}

	return task, nil
}