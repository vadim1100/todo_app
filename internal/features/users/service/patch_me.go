package users_service

import (
	"context"
	"fmt"

	core_domain "github.com/vadim1100/todo_app/internal/core/domain"
)

func (s *UsersService) PatchMe(ctx context.Context, id int, input PatchInput) (core_domain.User, error) {
	user, err := s.usersRepository.GetByID(ctx, id)
	if err != nil {
		return core_domain.User{}, err
	}

	if input.Username != nil {
		if err := validateUsername(*input.Username); err != nil{
			return core_domain.User{}, err
		}
		user.Username = *input.Username
	}

	if input.Password != nil {
		if err := validatePassword(*input.Password); err != nil {
			return core_domain.User{}, err
		}

		hash, err := s.hasher.Hash(*input.Password)
		if err != nil {
			return core_domain.User{}, fmt.Errorf("hash password: %w", err)
		}
		user.PasswordHash = hash
	}

	updatedUser, err := s.usersRepository.Update(ctx, user)
	if err != nil {
		return core_domain.User{}, err
	}

	return updatedUser, nil
}