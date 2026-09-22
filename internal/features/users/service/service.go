package users_service

import (
	"context"

	core_domain "github.com/vadim1100/todo_app/internal/core/domain"
	core_service_hash "github.com/vadim1100/todo_app/internal/core/service/hash"
	core_service_jwt "github.com/vadim1100/todo_app/internal/core/service/jwt"
)

type UsersService struct {
	usersRepository UsersRepository
	hasher core_service_hash.Hasher
	jwtManager *core_service_jwt.Manager
}

type UsersRepository interface{
	Create(
		ctx context.Context,
		user core_domain.User,
	) (core_domain.User, error)
	GetByUsername(
		ctx context.Context,
		username string,
	) (core_domain.User, error)
	GetByID(
		ctx context.Context,
		id int,
	) (core_domain.User, error)
	DeleteByID(
		ctx context.Context,
		id int,
	) error
	Update(
		ctx context.Context,
		user core_domain.User,
	) (core_domain.User, error)
}

func NewUsersService(
	usersRepository UsersRepository,
	hasher core_service_hash.Hasher,
	jwtManager *core_service_jwt.Manager,
) *UsersService{
	return &UsersService{
		usersRepository: usersRepository,
		hasher: hasher,
		jwtManager: jwtManager,
	}
}