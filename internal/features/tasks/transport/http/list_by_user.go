package tasks_http_transport

import (
	"net/http"

	core_errors "github.com/vadim1100/todo_app/internal/core/errors"
	core_logger "github.com/vadim1100/todo_app/internal/core/logger"
	core_http_middleware "github.com/vadim1100/todo_app/internal/core/transport/http/middleware"
	core_http_response "github.com/vadim1100/todo_app/internal/core/transport/http/response"
)

func (h *TasksHTTPHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, w)

	userID, ok := core_http_middleware.UserIDFromContext(ctx)
	if !ok {
		responseHandler.ErrorResponse(core_errors.ErrUnauthorized, "unauthorized")
		return
	}

	tasks, err := h.tasksService.ListByUser(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get tasks")
		return
	}
	
	responses := make([]TaskResponse, 0, len(tasks))
	for _, task := range tasks {
		responses = append(responses, NewTaskResponse(task))
	}

	responseHandler.JSONResponse(responses, http.StatusOK)
}