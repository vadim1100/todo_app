package core_domain

var (
	UninitializedVersion = -1
	UninitializedID = -1
)

type User struct{
	ID int
	Version int
	Username string
	PasswordHash string
}

func NewUser(
	id int,
	version int,
	username string,
	passwordHash string,
) User {
	return User{
		ID: id,
		Version: version,
		Username: username,
		PasswordHash: passwordHash,
	}
}

func NewUserUninitialized(
	username string,
	passwordHash string,
) User {
	return NewUser(
		UninitializedID,
		UninitializedVersion,
		username,
		passwordHash,
	)
}