package users_transport_http

type UserResponse struct {
    ID       int  `json:"id"`
    Username string `json:"username"`
}