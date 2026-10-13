package users_transport_http

type UserResponse struct {
    ID       int  `json:"id" example:"1"`
    Username string `json:"username" example:"vasya"`
}

type ErrorResponse struct {
	Error   string `json:"error" example:"unauthorized"`
	Message string `json:"message" example:"failed to login"`
}


func NewUserResponse(id int, username string) UserResponse{
    return UserResponse{
        ID: id,
        Username: username,
    }
}