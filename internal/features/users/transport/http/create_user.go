package users_transport_http

import (
	"encoding/json"
	"fmt"
	"net/http"

	core_logger "github.com/emp2ty0/golang-todoapp/internal/core/logger"
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
	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	log.Debug("invoce CreateUser handler")

	var request CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Println("Error:", err)
	}

	w.WriteHeader(http.StatusOK)
}
