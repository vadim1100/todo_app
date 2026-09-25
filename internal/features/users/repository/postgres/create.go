package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	core_domain "github.com/vadim1100/todo_app/internal/core/domain"
	core_errors "github.com/vadim1100/todo_app/internal/core/errors"
)

func (r *UsersRepository) Create(ctx context.Context, user core_domain.User) (core_domain.User, error) {
	query := `
		INSERT INTO todoapp.users (username, password_hash)
		VALUES ($1, $2)
		RETURNING id, version
	`

	row := r.pool.QueryRow(ctx, query, user.Username, user.PasswordHash)
	if err := row.Scan(&user.ID, &user.Version); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return core_domain.User{}, core_errors.ErrConflict
		}
		return core_domain.User{}, fmt.Errorf("create user: %w", err)
	}

	return user, nil
}