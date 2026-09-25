package tasks_http_transport

import (
	"context"
	"net/http"

	core_domain "github.com/vadim1100/todo_app/internal/core/domain"
	core_http_server "github.com/vadim1100/todo_app/internal/core/transport/http/server"
	tasks_service "github.com/vadim1100/todo_app/internal/features/tasks/service"
)

type TasksHTTPHandler struct {
	tasksService TasksService
}

type TasksService interface{
	Create(
		ctx context.Context,
		userID int,
		input tasks_service.CreateTaskInput,
	) (core_domain.Task, error)
}

func NewTasksHandler(tasksService TasksService) *TasksHTTPHandler {
	return &TasksHTTPHandler{
		tasksService: tasksService,
	}
}

func (h *TasksHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method: http.MethodPost,
			Path: "/tasks",
			Handler: h.Create,
			Protected: true,
		},
	}
}