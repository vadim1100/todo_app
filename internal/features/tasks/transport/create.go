package tasks_http_transport

import (
	"encoding/json"
	"net/http"
	core_errors "github.com/vadim1100/todo_app/internal/core/errors"
	core_logger "github.com/vadim1100/todo_app/internal/core/logger"
	core_http_middleware "github.com/vadim1100/todo_app/internal/core/transport/http/middleware"
	core_http_response "github.com/vadim1100/todo_app/internal/core/transport/http/response"
	tasks_service "github.com/vadim1100/todo_app/internal/features/tasks/service"
)

type CreateTaskRequest struct {
	Title string `json:"title"`
	Description *string `json:"description"`
}

func (h *TasksHTTPHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, w)

	userID, ok := core_http_middleware.UserIDFromContext(ctx)

	if !ok {
		responseHandler.ErrorResponse(core_errors.ErrUnauthorized, "unauthorized")
		return
	}

	var request CreateTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode create task http request")
		return
	}

	input := tasks_service.NewCreateTaskInput(request.Title, request.Description)

	task, err := h.tasksService.Create(ctx, userID, input)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create task")
		return
	}
	
	responseHandler.JSONResponse(NewTaskResponse(task), http.StatusCreated)
}