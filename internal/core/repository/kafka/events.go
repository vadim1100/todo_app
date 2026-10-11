package core_kafka

import "time"

const (
	EventTaskCreated = "task_created"
	EventTaskUpdated = "task_updated"
	EventTaskDeleted = "task_deleted"
)

type TaskEvent struct {
	Event     string    `json:"event"`
	UserID    int       `json:"user_id"`
	TaskID    int       `json:"task_id,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

func NewTaskEvent(event string, userID int, taskID int, timestamp time.Time) TaskEvent {
	return TaskEvent{
		Event:     event,
		UserID:    userID,
		TaskID:    taskID,
		Timestamp: timestamp,
	}
}
