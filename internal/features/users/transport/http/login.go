package users_transport_http

import (
	"encoding/json"
	"net/http"

	core_logger "github.com/vadim1100/todo_app/internal/core/logger"
	core_http_response "github.com/vadim1100/todo_app/internal/core/transport/http/response"
	users_service "github.com/vadim1100/todo_app/internal/features/users/service"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User UserResponse `json:"user"`
	Token string `json:"token"`
}

func (h *UsersHTTPHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)
	var request LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode login http request")
		return
	}

	input := users_service.NewAuthInput(request.Username, request.Password)
	output, err := h.usersService.Login(ctx, input)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to login")
		return
	}
	
	response := LoginResponse{
		User: UserResponse{
			ID: output.User.ID,
			Username: output.User.Username,
		},
		Token: output.Token,
	}

	responseHandler.JSONResponse(response, http.StatusOK)
}