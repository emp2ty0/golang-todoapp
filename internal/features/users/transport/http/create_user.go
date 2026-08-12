package users_transport_http

import (
	"net/http"

	"github.com/emp2ty0/golang-todoapp/internal/core/domain"
	core_logger "github.com/emp2ty0/golang-todoapp/internal/core/logger"
	core_http_request "github.com/emp2ty0/golang-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/emp2ty0/golang-todoapp/internal/core/transport/http/response"
)

type CreateUserRequest struct {
	Full_Name   string  `json:"full_name" validate:"required, min=3,max=100"`
	PhoneNumebr *string `json:"phoe_number" validate:"omitempty,min=10,max=15,startswith=+"`
}

type CreateUserResponse struct {
	ID          int    `json:"id"`
	Version     int    `json:"version"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
}

func (h *UserHTTPHandler) CreateUer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	log.Debug("invoce CreateUser handler")

	var request CreateUserRequest
	if err := core_http_request.DecodeAndvalidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failde to decode and valudate HTTP request")
		return
	}

	userDomain := domainFromDTO(request)

	userDomain, err := h.userService.CreateUser(ctx, userDomain)

	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create user")
		return
	}

	response := dtoFromDomain(userDomain)

	responseHandler.JSONResponse(response, http.StatusCreated)

}

func domainFromDTO(dto CreateUserRequest) domain.User {
	return domain.NewUserUnitialized(dto.Full_Name, dto.PhoneNumebr)
}

func dtoFromDomain(user domain.User) CreateUserResponse {
	return CreateUserResponse{
		ID:          user.Id,
		Version:     user.Version,
		FullName:    user.FullName,
		PhoneNumber: *user.PhoneNumber,
	}
}
