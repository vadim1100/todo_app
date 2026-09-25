package core_domain

import "time"

type Task struct {
	ID          int
	Version     int
	UserID      int
	Title       string
	Description *string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

func NewTaskUninitialized(
	userID int,
	title string,
	description *string,
) Task {
	return Task{
		ID:          UninitializedID,
		Version:     UninitializedVersion,
		UserID:      userID,
		Title:       title,
		Description: description,
	}
}
