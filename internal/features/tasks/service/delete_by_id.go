package tasks_service

import "context"

func (s *TasksService) DeleteByTaskID(ctx context.Context,
	taskID int,
	userID int,
) error {
	return s.tasksRepository.DeleteByTaskID(ctx, taskID, userID)
}