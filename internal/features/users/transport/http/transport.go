package users_transport_http

import core_http_server "github.com/vadim1100/todo_app/internal/core/transport/http/server"

type UsersHTTPHandler struct {
	usersService UsersService
}

type UsersService interface{

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
		
	}
}