package users_service

import (
	"context"
	"fmt"

	core_errors "github.com/vadim1100/todo_app/internal/core/errors"
)

func (s *UsersService) Login(ctx context.Context, input AuthInput) (AuthOutput, error){
	user, err := s.usersRepository.GetByUsername(ctx, input.Username)

	if err != nil {
		return AuthOutput{}, core_errors.ErrNotFound
	}

	if err := s.hasher.Compare(user.PasswordHash, input.Password); err != nil {
		return AuthOutput{}, core_errors.ErrUnauthorized
	}

	token, err := s.jwtManager.Generate(user.ID)
	if err != nil {
		return AuthOutput{}, fmt.Errorf("failed to generate jwt token")
	}

	return AuthOutput{
		User: user,
		Token: token,
	}, nil
}