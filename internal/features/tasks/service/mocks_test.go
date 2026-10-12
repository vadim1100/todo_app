package tasks_service_test

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"

	core_domain "github.com/vadim1100/todo_app/internal/core/domain"
)

type mockTasksRepo struct {
	mock.Mock
}

type mockCache struct {
	mock.Mock
}

func (m *mockTasksRepo) Create(ctx context.Context, task core_domain.Task) (core_domain.Task, error) {
	args := m.Called(ctx, task)
	return args.Get(0).(core_domain.Task), args.Error(1)
}

func (m *mockTasksRepo) GetByTaskID(ctx context.Context, taskID, userID int) (core_domain.Task, error) {
	args := m.Called(ctx, taskID, userID)
	return args.Get(0).(core_domain.Task), args.Error(1)
}

func (m *mockTasksRepo) ListByUser(ctx context.Context, userID int) ([]core_domain.Task, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]core_domain.Task), args.Error(1)
}

func (m *mockTasksRepo) DeleteByTaskID(ctx context.Context, taskID, userID int) error {
	args := m.Called(ctx, taskID, userID)
	return args.Error(0)
}

func (m *mockTasksRepo) Update(ctx context.Context, task core_domain.Task) (core_domain.Task, error) {
	args := m.Called(ctx, task)
	return args.Get(0).(core_domain.Task), args.Error(1)
}


func (m *mockCache) Get(ctx context.Context, key string) ([]byte, error) {
	args := m.Called(ctx, key)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]byte), args.Error(1)
}

func (m *mockCache) Set(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	args := m.Called(ctx, key, value, ttl)
	return args.Error(0)
}

func (m *mockCache) Delete(ctx context.Context, keys ...string) error {
	args := m.Called(ctx, keys)
	return args.Error(0)
}

type mockProducer struct {
	mock.Mock
}

func (m *mockProducer) Produce(ctx context.Context, key string, event any) error {
	args := m.Called(ctx, key, event)
	return args.Error(0)
}