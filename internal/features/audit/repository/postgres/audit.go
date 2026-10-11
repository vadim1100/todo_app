package audit_postgres_repository

import (
	"context"
	"fmt"
	"time"

	core_postgres_pool "github.com/vadim1100/todo_app/internal/core/repository/postgres/pool"
)

type AuditRepository struct {
	pool core_postgres_pool.Pool
}

func NewAuditRepository(pool core_postgres_pool.Pool) *AuditRepository {
	return &AuditRepository{
		pool: pool,
	}
}

func (r *AuditRepository) Insert(
	ctx context.Context,
	event string,
	userID int,
	taskID *int,
	payload []byte,
) error {
	query := `
	INSERT INTO todoapp.audit_log (event, user_id, task_id, created_at, payload)
	VALUES ($1, $2, $3, $4, $5)
	`

	if _, err := r.pool.Exec(ctx, query, event, userID, taskID, time.Now().UTC(), payload); err != nil {
		return fmt.Errorf("insert audit log: %w", err)
	}
	return nil
}