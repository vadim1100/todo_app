package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	core_domain "github.com/vadim1100/todo_app/internal/core/domain"
	core_errors "github.com/vadim1100/todo_app/internal/core/errors"
)

func (r *TasksRepository) GetByTaskID(
	ctx context.Context,
	taskID int,
	userID int,
) (core_domain.Task, error) {
	query := `
	SELECT id, version, user_id, title, description, completed, created_at, completed_at FROM todoapp.tasks
	WHERE id = $1 AND user_id = $2
	`

	var task core_domain.Task

	row := r.pool.QueryRow(ctx, query, taskID, userID)
	if err := row.Scan(
		&task.ID,
		&task.Version,
		&task.UserID,
		&task.Title,
		&task.Description,
		&task.Completed,
		&task.CreatedAt,
		&task.CompletedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return core_domain.Task{}, core_errors.ErrNotFound
		}
		return core_domain.Task{}, fmt.Errorf("get task by id: %w", err)
	}

	return task, nil
}