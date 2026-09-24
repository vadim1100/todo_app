package users_transport_http

type UserResponse struct {
    ID       int  `json:"id"`
    Username string `json:"username"`
}

func NewUserResponse(id int, username string) UserResponse{
    return UserResponse{
        ID: id,
        Username: username,
    }
}