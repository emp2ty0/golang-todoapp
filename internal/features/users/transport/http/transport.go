package users_transport_http

import (
	"net/http"

	core_http_server "github.com/emp2ty0/golang-todoapp/internal/core/transport/http/server"
)

type UserHTTPHandler struct {
	userService UserService
}

type UserService interface {
}

func NewUserHTTPHandler(userSerice UserService) *UserHTTPHandler {
	return &UserHTTPHandler{
		userService: userSerice,
	}
}

func (h *UserHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUer,
		},
	}
}
