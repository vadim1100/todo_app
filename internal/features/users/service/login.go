package users_service

import (
	"context"
	"errors"
	"fmt"

	core_errors "github.com/vadim1100/todo_app/internal/core/errors"
)

func (s *UsersService) Login(ctx context.Context, input AuthInput) (AuthOutput, error){
	user, err := s.usersRepository.GetByUsername(ctx, input.Username)

	if err != nil {
		if errors.Is(err, core_errors.ErrNotFound) {
			return AuthOutput{}, core_errors.ErrUnauthorized
		}
		return AuthOutput{}, fmt.Errorf("get user by username: %w", err)
	}

	if err := s.hasher.Compare(user.PasswordHash, input.Password); err != nil {
		return AuthOutput{}, core_errors.ErrUnauthorized
	}

	token, err := s.jwtManager.Generate(user.ID)
	if err != nil {
		return AuthOutput{}, fmt.Errorf("failed to generate jwt token: %w", err)
	}

	return NewAuthOutput(
		user,
		token,
	), nil
}