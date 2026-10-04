package users_transport_http

import (
	"net/http"

	core_errors "github.com/vadim1100/todo_app/internal/core/errors"
	core_logger "github.com/vadim1100/todo_app/internal/core/logger"
	core_http_middleware "github.com/vadim1100/todo_app/internal/core/transport/http/middleware"
	core_http_response "github.com/vadim1100/todo_app/internal/core/transport/http/response"
)

func (h *UsersHTTPHandler) Logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, w)

	token, ok := core_http_middleware.TokenFromContext(ctx)
	if !ok {
		responseHandler.ErrorResponse(core_errors.ErrUnauthorized, "missing token")
		return
	}

	if err := h.usersService.Logout(ctx, token); err != nil {
		responseHandler.ErrorResponse(err, "failed to logout")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}