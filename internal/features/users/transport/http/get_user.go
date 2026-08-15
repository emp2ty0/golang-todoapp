package users_transport_http

import (
	"fmt"
	"net/http"

	core_errors "github.com/emp2ty0/golang-todoapp/internal/core/errors"
	core_logger "github.com/emp2ty0/golang-todoapp/internal/core/logger"
	core_http_response "github.com/emp2ty0/golang-todoapp/internal/core/transport/http/response"
	core_http_utils "github.com/emp2ty0/golang-todoapp/internal/core/transport/http/utils"
)

func (u *UserHTTPHandler) GetUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)
	userId, err := getUserId(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user id query params")
	}

	user, err := u.userService.GetUser(ctx, userId)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user")
	}

	responseHandler.JSONResponse(user, http.StatusOK)
}

func getUserId(r *http.Request) (*int, error) {
	userId, err := core_http_utils.GetIntQueryParams(r, "id")
	if err != nil {
		return nil, fmt.Errorf("get user id: %w", core_errors.ErrInvalidArgument)
	}

	return userId, nil
}
