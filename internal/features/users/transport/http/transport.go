package users_transport_http

import (
	"context"
	"net/http"

	core_http_server "github.com/vadim1100/todo_app/internal/core/transport/http/server"
	users_service "github.com/vadim1100/todo_app/internal/features/users/service"
)

type UsersHTTPHandler struct {
	usersService UsersService
}

type UsersService interface{
	Register(
		ctx context.Context,
		in users_service.AuthInput,
	) (users_service.AuthOutput, error)
	Login(
		ctx context.Context,
		in users_service.AuthInput,
	) (users_service.AuthOutput, error)
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
	}
}