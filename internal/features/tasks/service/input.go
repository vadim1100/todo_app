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