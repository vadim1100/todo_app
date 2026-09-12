package core_domain

type User struct{
	ID int
	Version int
	Username string
	PasswordHash string
}