package users_transport_http

import (
	"context"
	"net/http"

	core_domain "github.com/vadim1100/todo_app/internal/core/domain"
	core_http_server "github.com/vadim1100/todo_app/internal/core/transport/http/server"
	users_service "github.com/vadim1100/todo_app/internal/features/users/service"
)

type UsersHTTPHandler struct {
	usersService UsersService
}

type UsersService interface{
	Register(
		ctx context.Context,
		input users_service.AuthInput,
	) (users_service.AuthOutput, error)
	Login(
		ctx context.Context,
		input users_service.AuthInput,
	) (users_service.AuthOutput, error)
	GetMe(
		ctx context.Context,
		id int,
	) (core_domain.User, error)
	DeleteMe(
		ctx context.Context,
		id int,
	) error
	PatchMe(
		ctx context.Context,
		id int,
		input users_service.PatchInput,
	) (core_domain.User, error)
}

func NewUsersHTTPHandler(
	usersService UsersService,
) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		usersService: usersService,
	}
}

func (h *UsersHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method: http.MethodPost,
			Path: "/users/register",
			Handler: h.Register,
		},
		{
			Method: http.MethodPost,
			Path: "/users/login",
			Handler: h.Login,
		},
		{
			Method: http.MethodGet,
			Path: "/users/me",
			Handler: h.GetMe,
			Protected: true,
		},
		{
			Method: http.MethodDelete,
			Path: "/users/me",
			Handler: h.DeleteMe,
			Protected: true,
		},
		{
			Method: http.MethodPatch,
			Path: "/users/me",
			Handler: h.PatchMe,
			Protected: true,
		},
	}
}