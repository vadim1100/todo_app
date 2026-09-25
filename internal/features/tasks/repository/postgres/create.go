package tasks_postgres_repository

import (
	"context"
	"fmt"

	core_domain "github.com/vadim1100/todo_app/internal/core/domain"
)

func (r *TasksRepository) Create(ctx context.Context, task core_domain.Task) (core_domain.Task, error) {
	query := `
	INSERT INTO todoapp.tasks (user_id, title, description, completed)
	VALUES ($1, $2, $3, $4)
	RETURNING id, version, created_at
	`

	row := r.pool.QueryRow(ctx, query, task.UserID, task.Title, task.Description, task.Completed)
	if err := row.Scan(&task.ID, &task.Version, &task.CreatedAt); err != nil {
		return core_domain.Task{}, fmt.Errorf("create task: %w", err)
	}
	return task, nil
}