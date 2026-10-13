package users_transport_http

import (
	"encoding/json"
	"net/http"

	core_errors "github.com/vadim1100/todo_app/internal/core/errors"
	core_logger "github.com/vadim1100/todo_app/internal/core/logger"
	core_http_middleware "github.com/vadim1100/todo_app/internal/core/transport/http/middleware"
	core_http_response "github.com/vadim1100/todo_app/internal/core/transport/http/response"
	users_service "github.com/vadim1100/todo_app/internal/features/users/service"
)

type PatchRequest struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}

// PatchMe godoc
// @Summary      Обновить профиль
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body PatchRequest true "Новые данные"
// @Success      200 {object} UserResponse
// @Failure      400 {object} ErrorResponse "невалидные данные"
// @Failure      401 {object} ErrorResponse "нет токена"
// @Failure      404 {object} ErrorResponse "пользователь не найден"
// @Failure      409 {object} ErrorResponse "username занят"
// @Failure      500 {object} ErrorResponse "внутренняя ошибка"
// @Router       /users/me [patch]
func (h *UsersHTTPHandler) PatchMe(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, w)

	userID, ok := core_http_middleware.UserIDFromContext(ctx)
	if !ok {
		responseHandler.ErrorResponse(core_errors.ErrUnauthorized, "unauthorized")
		return
	}

	var request PatchRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode patch http request")
		return
	}

	user, err := h.usersService.PatchMe(ctx, userID, users_service.NewPatchInput(
		request.Username,
		request.Password,
	))
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to update user")
		return
	}

	responseHandler.JSONResponse(
		NewUserResponse(user.ID, user.Username),
		http.StatusOK,
	)
}