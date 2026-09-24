package users_repository

import (
	"context"
	"fmt"
	core_errors "github.com/vadim1100/todo_app/internal/core/errors"
)

func (r *UsersRepository) DeleteByID(ctx context.Context, id int) error{
	query := `
	DELETE FROM todoapp.users WHERE id = $1
	`

	commandTag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("delete user: %w", err)
	}

	if commandTag.RowsAffected() == 0{
		return core_errors.ErrNotFound
	}

	return nil
}