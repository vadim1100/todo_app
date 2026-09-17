package users_transport_http

import (
	"encoding/json"
	"net/http"

	core_logger "github.com/vadim1100/todo_app/internal/core/logger"
	core_http_response "github.com/vadim1100/todo_app/internal/core/transport/http/response"
	users_service "github.com/vadim1100/todo_app/internal/features/users/service"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResponse struct{
	User UserResponse `json:"user"`
	Token string `json:"token"`
}

func (h *UsersHTTPHandler) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)
	var request RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode http request")
		return
	}

	input := users_service.NewRegisterInput(request.Username, request.Password)
	output, err := h.usersService.Register(ctx, input)

	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create user")
		return
	}

	response := RegisterResponse{
		User: UserResponse{
			ID: output.User.ID,
			Username: output.User.Username,
		},
		Token: output.Token,
	}

	responseHandler.JSONResponse(response, http.StatusCreated)
}
