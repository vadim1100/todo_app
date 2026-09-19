package users_service

type AuthInput struct{
	Username string
	Password string
}

func NewAuthInput(username string, password string) AuthInput{
	return AuthInput{
		Username: username,
		Password: password,
	}
}