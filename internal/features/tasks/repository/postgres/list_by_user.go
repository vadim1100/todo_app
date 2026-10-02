package tasks_postgres_repository

import (
	"context"
	"fmt"

	core_domain "github.com/vadim1100/todo_app/internal/core/domain"
)

func (r *TasksRepository) ListByUser(
	ctx context.Context,
	userID int,
) ([]core_domain.Task, error) {
	query := `
	SELECT id, version, user_id, title, description, completed, created_at, completed_at FROM todoapp.tasks
	WHERE user_id = $1
	ORDER BY created_at DESC
	`

	rows, err := r.pool.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("list tasks by user: %w", err)
	}
	defer rows.Close()

	var tasks []core_domain.Task

	for rows.Next() {
		var task core_domain.Task
		if err := rows.Scan(
			&task.ID,
			&task.Version,
			&task.UserID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.CreatedAt,
			&task.CompletedAt,
		); err != nil {
			return nil, fmt.Errorf("scan tasks: %w", err)
		}
		tasks = append(tasks, task)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return tasks, nil
}
