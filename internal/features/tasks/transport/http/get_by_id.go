package tasks_http_transport

import (
	"net/http"
	"strconv"

	core_errors "github.com/vadim1100/todo_app/internal/core/errors"
	core_logger "github.com/vadim1100/todo_app/internal/core/logger"
	core_http_middleware "github.com/vadim1100/todo_app/internal/core/transport/http/middleware"
	core_http_response "github.com/vadim1100/todo_app/internal/core/transport/http/response"
)

func (h *TasksHTTPHandler) GetByTaskID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, w)

	userID, ok := core_http_middleware.UserIDFromContext(ctx)
	if !ok {
		responseHandler.ErrorResponse(core_errors.ErrUnauthorized, "unauthorized")
		return
	}

	taskID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		responseHandler.ErrorResponse(core_errors.ErrInvalidArgument, "invalid task id")
		return
	}
	
	task, err := h.tasksService.GetByTaskID(ctx, taskID, userID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get task")
		return
	}

	responseHandler.JSONResponse(NewTaskResponse(task), http.StatusOK)
}