package tasks_service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	core_domain "github.com/vadim1100/todo_app/internal/core/domain"
	tasks_service "github.com/vadim1100/todo_app/internal/features/tasks/service"
)

func TestCreate_OK(t *testing.T) {
	ctx := context.Background()
	userID := 1
	input := tasks_service.CreateTaskInput{Title: "Купить хлеб"}

	expectedTask := core_domain.Task{
		ID:     42,
		UserID: userID,
		Title:  "Купить хлеб",
	}

	repo := new(mockTasksRepo)
	cache := new(mockCache)
	producer := new(mockProducer)

	repo.On("Create", ctx, mock.Anything).Return(expectedTask, nil).Once()
	cache.On("Delete", ctx, mock.Anything).Return(nil).Once()
	producer.On("Produce", ctx, "1", mock.Anything).Return(nil).Once()

	svc := tasks_service.NewTasksService(repo, cache, producer)

	
	got, err := svc.Create(ctx, userID, input)

	
	assert.NoError(t, err)
	assert.Equal(t, 42, got.ID)
	assert.Equal(t, "Купить хлеб", got.Title)

	repo.AssertExpectations(t)
	cache.AssertExpectations(t)
	producer.AssertExpectations(t)
}

func TestCreate_EmptyTitle_ReturnsError(t *testing.T) {
	ctx := context.Background()
	input := tasks_service.CreateTaskInput{Title: ""}

	repo := new(mockTasksRepo)
	cache := new(mockCache)
	producer := new(mockProducer)

	svc := tasks_service.NewTasksService(repo, cache, producer)


	_, err := svc.Create(ctx, 1, input)


	assert.Error(t, err)


	repo.AssertNotCalled(t, "Create", mock.Anything, mock.Anything)
	cache.AssertNotCalled(t, "Delete", mock.Anything, mock.Anything)
	producer.AssertNotCalled(t, "Produce", mock.Anything, mock.Anything, mock.Anything)
}

func TestCreate_RepoError_ReturnsError(t *testing.T) {
	ctx := context.Background()
	input := tasks_service.CreateTaskInput{Title: "Купить хлеб"}
	repoErr := errors.New("db is down")

	repo := new(mockTasksRepo)
	cache := new(mockCache)
	producer := new(mockProducer)

	repo.On("Create", ctx, mock.Anything).Return(core_domain.Task{}, repoErr).Once()

	svc := tasks_service.NewTasksService(repo, cache, producer)


	_, err := svc.Create(ctx, 1, input)

	
	assert.Error(t, err)
	assert.ErrorContains(t, err, "db is down")

	cache.AssertNotCalled(t, "Delete", mock.Anything, mock.Anything)
	producer.AssertNotCalled(t, "Produce", mock.Anything, mock.Anything, mock.Anything)
}

func TestCreate_ProducerError_TaskStillCreated(t *testing.T) {
	
	ctx := context.Background()
	userID := 1
	input := tasks_service.CreateTaskInput{Title: "Купить хлеб"}
	producerErr := errors.New("kafka is down")

	expectedTask := core_domain.Task{ID: 42, UserID: userID, Title: "Купить хлеб"}

	repo := new(mockTasksRepo)
	cache := new(mockCache)
	producer := new(mockProducer)

	repo.On("Create", ctx, mock.Anything).Return(expectedTask, nil).Once()
	cache.On("Delete", ctx, mock.Anything).Return(nil).Once()
	producer.On("Produce", ctx, "1", mock.Anything).Return(producerErr).Once()

	svc := tasks_service.NewTasksService(repo, cache, producer)


	got, err := svc.Create(ctx, userID, input)


	assert.NoError(t, err)
	assert.Equal(t, 42, got.ID)

	repo.AssertExpectations(t)
	cache.AssertExpectations(t)
	producer.AssertExpectations(t)
}