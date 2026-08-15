package users_transport_http

import (
	"context"
	"net/http"

	"github.com/emp2ty0/golang-todoapp/internal/core/domain"
	core_http_server "github.com/emp2ty0/golang-todoapp/internal/core/transport/http/server"
)

type UserHTTPHandler struct {
	userService UserService
}

type UserService interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)

	GetUsers(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.User, error)

	GetUser(
		ctx context.Context,
		id *int,
	) (domain.User, error)
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
		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: h.GetUsers,
		},
		{
			Method:  http.MethodGet,
			Path:    "/user",
			Handler: h.GetUser,
		},
	}
}
