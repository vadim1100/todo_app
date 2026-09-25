package users_service

import (
	"context"
	"fmt"

	core_domain "github.com/vadim1100/todo_app/internal/core/domain"
	core_errors "github.com/vadim1100/todo_app/internal/core/errors"
)

func (s *UsersService) Register(ctx context.Context, input AuthInput) (AuthOutput, error) {
	if err := validateUsername(input.Username); err != nil {
		return AuthOutput{}, fmt.Errorf("validate username %w", err)
	}
	if err := validatePassword(input.Password); err != nil {
		return AuthOutput{}, fmt.Errorf("validate password %w", err)
	}

	hash, err := s.hasher.Hash(input.Password)

	if err != nil {
		return AuthOutput{}, fmt.Errorf("hash password %w", err)
	}

	user := core_domain.NewUserUninitialized(input.Username, hash)

	createdUser, err := s.usersRepository.Create(ctx, user)
	if err != nil {
		return AuthOutput{}, fmt.Errorf("create user in repository: %w", err)
	}

	token, err := s.jwtManager.Generate(createdUser.ID)
	if err != nil {
		return AuthOutput{}, fmt.Errorf("generate jwt token: %w", err)
	}
	return AuthOutput{
		User: createdUser,
		Token: token,
	}, nil
}

func validateUsername(username string) error {
	usernameLen := len([]rune(username))
    if usernameLen < 1 || usernameLen > 100 {
        return fmt.Errorf("%w: username must be 1-100 characters", core_errors.ErrInvalidArgument)
    }
    return nil
}

func validatePassword(password string) error {
    if len(password) < 8 {
        return fmt.Errorf("%w: password must be at least 8 characters", core_errors.ErrInvalidArgument)
    }
    if len(password) > 72 {
        return fmt.Errorf("%w: password must be at most 72 characters", core_errors.ErrInvalidArgument)
    }
    return nil
}