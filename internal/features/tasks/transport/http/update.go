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
	Title       *string `json:"title,omitempty" example:"Измененный заголовок"`
	Description *string `json:"description,omitempty"`
	Completed   *bool   `json:"completed,omitempty" example:"true"`
}

// UpdateTask godoc
// @Summary      Обновить задачу
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID задачи"
// @Param        request body UpdateTaskRequest true "Тело"
// @Success      200 {object} TaskResponse
// @Failure      400 {object} ErrorResponse "невалидные данные"
// @Failure      401 {object} ErrorResponse "нет токена"
// @Failure      404 {object} ErrorResponse "задача не найдена"
// @Failure      500 {object} ErrorResponse "внутренняя ошибка"
// @Router       /tasks/{id} [patch]
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
