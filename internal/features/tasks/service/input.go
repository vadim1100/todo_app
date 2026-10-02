package tasks_service

type CreateTaskInput struct {
	Title string
	Description *string
}

func NewCreateTaskInput(
	title string,
	description *string,
) CreateTaskInput {
	return CreateTaskInput{
		Title: title,
		Description: description,
	}
}

type UpdateTaskInput struct {
	Title *string
	Description *string
	Completed *bool
}

func NewUpdateTaskInput(
	title *string,
	description *string,
	completed *bool,
) UpdateTaskInput{
	return UpdateTaskInput{
		Title: title,
		Description: description,
		Completed: completed,
	}
}