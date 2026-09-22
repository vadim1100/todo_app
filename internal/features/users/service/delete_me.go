package users_service

import (
	"context"
)

func (s *UsersService) DeleteMe(ctx context.Context, id int) error {
	return s.usersRepository.DeleteByID(ctx, id)
}