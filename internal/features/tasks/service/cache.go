package tasks_service

import "time"

const (
	tasksListTTL = 60 * time.Second
	taskItemTTL = 5 * time.Minute
)