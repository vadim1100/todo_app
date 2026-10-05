package core_redis

import "strconv"

const (
	blacklistPrefix = "blacklist:"

	tasksListPrefix = "tasks:v1:user:"
	taskItemPrefix = "tasks:v1:item:"
)

func NewBlacklistKey(token string) string {
	return blacklistPrefix + token
}

func NewTasksListKey(userID int) string {
	return tasksListPrefix + strconv.Itoa(userID) + ":list"
}

func NewTaskItemKey(taskID int) string {
	return taskItemPrefix + strconv.Itoa(taskID)
}