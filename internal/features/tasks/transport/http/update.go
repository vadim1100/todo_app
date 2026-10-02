package tasks_http_transport

import (
	"encoding/json"
	"net/http"
	"strconv"

	core_errors "github.com/vadim1100/todo_app/internal/core/errors"
	core_logger "github.com/vadim1100/todo_app/internal/core/logger"
	core_http_middleware "github.com/vadim1100/todo_app/internal/core/transport/http/middleware"
	core_http_response "github.com/vadim1100/todo_app/internal/core/transport/http/response"
	tasks_service "github.com/vadim1100/todo_app/internal/features/tasks/service"
)

type UpdateTaskRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Completed   *bool   `json:"completed"`
}

func (h *TasksHTTPHandler) Update(w http.ResponseWriter, r *http.Request) {
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

	var request UpdateTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode update task http request")
		return
	}

	taskInput := tasks_service.NewUpdateTaskInput(request.Title, request.Description, request.Completed)

	task, err := h.tasksService.Update(ctx, taskID, userID, taskInput)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to update task")
		return
	}

	responseHandler.JSONResponse(NewTaskResponse(task), http.StatusOK)
}
