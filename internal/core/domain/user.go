package core_domain

type User struct{
	ID int
	Version int
	Username string
	PasswordHash string
}

func NewUserUninitialized(
	username string,
	passwordHash string,
) User {
	return User{
		ID: UninitializedID,
		Version: UninitializedVersion,
		Username: username,
		PasswordHash: passwordHash,
	}
}