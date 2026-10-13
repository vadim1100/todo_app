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

// Register godoc
// @Summary      Регистрация
// @Description  Создаёт пользователя, хеширует пароль (bcrypt)
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request body RegisterRequest true "Данные регистрации"
// @Success      201 {object} RegisterResponse
// @Failure      400 {object} ErrorResponse "невалидные данные"
// @Failure      409 {object} ErrorResponse "пользователь уже существует"
// @Failure      500 {object} ErrorResponse "внутренняя ошибка"
// @Router       /users/register [post]
func (h *UsersHTTPHandler) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)
	var request RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode register http request")
		return
	}

	input := users_service.NewAuthInput(request.Username, request.Password)
	output, err := h.usersService.Register(ctx, input)

	if err != nil {
		responseHandler.ErrorResponse(err, "failed to register user")
		return
	}

	response := NewRegisterResponse(
		NewUserResponse(output.User.ID, output.User.Username),
		output.Token,
	)

	responseHandler.JSONResponse(response, http.StatusCreated)
}

func NewRegisterResponse(user UserResponse, token string) RegisterResponse {
	return RegisterResponse{
		User: user,
		Token: token,
	}
}