package users_service

import (
	core_domain "github.com/vadim1100/todo_app/internal/core/domain"
)

type AuthOutput struct {
	User core_domain.User
	Token string
}