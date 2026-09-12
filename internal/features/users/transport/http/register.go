package users_transport_http

import (
	"encoding/json"
	"net/http"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *UsersHTTPHandler) Register(w http.ResponseWriter, r *http.Request) {
	var request RegisterRequest
	json.NewDecoder(r.Body).Decode(&request)
}