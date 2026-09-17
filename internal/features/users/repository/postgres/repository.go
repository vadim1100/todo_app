package users_repository

import (
	core_postgres_pool "github.com/vadim1100/todo_app/internal/core/repository/pool"
)

type UsersRepository struct {
	pool core_postgres_pool.Pool
}

func NewUsersRepository(
	pool core_postgres_pool.Pool,
) *UsersRepository {
	return &UsersRepository{
		pool: pool,
	}
}