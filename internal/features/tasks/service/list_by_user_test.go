package tasks_service_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	core_domain "github.com/vadim1100/todo_app/internal/core/domain"
	tasks_service "github.com/vadim1100/todo_app/internal/features/tasks/service"
)

func TestListByUser_CacheHit(t *testing.T) {
	ctx := context.Background()
	userID := 1

	cachedTasks := []core_domain.Task{
		{ID: 1, UserID: userID, Title: "Из кэша 1"},
		{ID: 2, UserID: userID, Title: "Из кэша 2"},
	}
	cachedJSON, _ := json.Marshal(cachedTasks)

	repo := new(mockTasksRepo)
	cache := new(mockCache)
	producer := new(mockProducer)

	cache.On("Get", ctx, mock.Anything).Return(cachedJSON, nil).Once()

	svc := tasks_service.NewTasksService(repo, cache, producer, testLogger())


	got, err := svc.ListByUser(ctx, userID)


	assert.NoError(t, err)
	assert.Len(t, got, 2)
	assert.Equal(t, "Из кэша 1", got[0].Title)


	repo.AssertNotCalled(t, "ListByUser", mock.Anything, mock.Anything)
	cache.AssertExpectations(t)
}

func TestListByUser_CacheMiss(t *testing.T) {
	ctx := context.Background()
	userID := 1

	fromDB := []core_domain.Task{
		{ID: 1, UserID: userID, Title: "Из БД 1"},
		{ID: 2, UserID: userID, Title: "Из БД 2"},
	}

	repo := new(mockTasksRepo)
	cache := new(mockCache)
	producer := new(mockProducer)


	cache.On("Get", ctx, mock.Anything).Return(nil, errors.New("cache miss")).Once()
	
	repo.On("ListByUser", ctx, userID).Return(fromDB, nil).Once()

	cache.On("Set", ctx, mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

	svc := tasks_service.NewTasksService(repo, cache, producer, testLogger())


	got, err := svc.ListByUser(ctx, userID)


	assert.NoError(t, err)
	assert.Len(t, got, 2)
	assert.Equal(t, "Из БД 1", got[0].Title)

	repo.AssertExpectations(t)
	cache.AssertExpectations(t)
}

func TestListByUser_RepoError(t *testing.T) {
	ctx := context.Background()
	userID := 1
	repoErr := errors.New("db is down")

	repo := new(mockTasksRepo)
	cache := new(mockCache)
	producer := new(mockProducer)

	cache.On("Get", ctx, mock.Anything).Return(nil, errors.New("cache miss")).Once()
	repo.On("ListByUser", ctx, userID).Return(nil, repoErr).Once()

	svc := tasks_service.NewTasksService(repo, cache, producer, testLogger())


	_, err := svc.ListByUser(ctx, userID)


	assert.Error(t, err)
	assert.ErrorContains(t, err, "db is down")


	cache.AssertNotCalled(t, "Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
	repo.AssertExpectations(t)
}