package users_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	core_domain "github.com/vadim1100/todo_app/internal/core/domain"
	core_errors "github.com/vadim1100/todo_app/internal/core/errors"
)

func (r *UsersRepository) GetByUsername(ctx context.Context, username string) (core_domain.User, error) {
	query := `
	SELECT id, version, username, password_hash FROM todoapp.users
	WHERE username = $1
	`

	var user core_domain.User

	row := r.pool.QueryRow(ctx, query, username)
	if err := row.Scan(&user.ID, &user.Version, &user.Username, &user.PasswordHash); err != nil{
		if errors.Is(err, pgx.ErrNoRows) {
			return core_domain.User{}, core_errors.ErrNotFound
		}
		return core_domain.User{}, fmt.Errorf("get user by username: %w", err)
	}
	
	return user, nil
}
