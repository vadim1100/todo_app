package tasks_service

import (
	"context"
	"time"

	core_domain "github.com/vadim1100/todo_app/internal/core/domain"
	core_logger "github.com/vadim1100/todo_app/internal/core/logger"
)

type TasksService struct {
	tasksRepository TasksRepository
	cache           Cache
	producer        EventProducer
	logger *core_logger.Logger
}

type TasksRepository interface {
	Create(
		ctx context.Context,
		task core_domain.Task,
	) (core_domain.Task, error)
	GetByTaskID(
		ctx context.Context,
		taskID int,
		userID int,
	) (core_domain.Task, error)
	ListByUser(
		ctx context.Context,
		userID int,
	) ([]core_domain.Task, error)
	DeleteByTaskID(
		ctx context.Context,
		taskID int,
		userID int,
	) error
	Update(
		ctx context.Context,
		task core_domain.Task,
	) (core_domain.Task, error)
}

type Cache interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, value []byte, ttl time.Duration) error
	Delete(ctx context.Context, keys ...string) error
}

type EventProducer interface {
	Produce(ctx context.Context, key string, event any) error
}

func NewTasksService(
	tasksRepository TasksRepository,
	cache Cache,
	producer EventProducer,
	logger *core_logger.Logger,
) *TasksService {
	return &TasksService{
		tasksRepository: tasksRepository,
		cache:           cache,
		producer:        producer,
		logger: logger,
	}
}
