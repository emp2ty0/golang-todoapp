package users_transport_http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateUserRequest struct {
	Full_Name   string  `json:"full_name"`
	PhoneNumebr *string `json:"phoe_number"`
}

type CreateUserResponse struct {
	ID          int    `json:"id"`
	Version     int    `json:"version"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
}

func (*UserHTTPHandler) CreateUer(w http.ResponseWriter, r *http.Request) {
	var request CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Println("Error:", err)
	}
}
