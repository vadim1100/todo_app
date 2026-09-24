package users_service

import (
	"context"

	core_domain "github.com/vadim1100/todo_app/internal/core/domain"
)

func (s *UsersService) GetMe(ctx context.Context, id int) (core_domain.User, error){
	return s.usersRepository.GetByID(ctx, id)
}