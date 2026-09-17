package users_service

type RegisterInput struct{
	Username string
	Password string
}

func NewRegisterInput(username string, password string) RegisterInput{
	return RegisterInput{
		Username: username,
		Password: password,
	}
}