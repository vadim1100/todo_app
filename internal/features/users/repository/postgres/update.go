package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	core_domain "github.com/vadim1100/todo_app/internal/core/domain"
	core_errors "github.com/vadim1100/todo_app/internal/core/errors"
)

func (r *UsersRepository) Update(ctx context.Context, user core_domain.User) (core_domain.User, error) {
	query := `
		UPDATE todoapp.users
		SET username = $1,
		password_hash = $2,
		version = version + 1
		WHERE id = $3 AND version = $4
		RETURNING version
	`

	row := r.pool.QueryRow(ctx, query, user.Username, user.PasswordHash, user.ID, user.Version)
	if err := row.Scan(&user.Version); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return core_domain.User{}, core_errors.ErrConflict
		}
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return core_domain.User{}, core_errors.ErrConflict
		}
		return core_domain.User{}, fmt.Errorf("update user: %w", err)
	}
	return user, nil
}