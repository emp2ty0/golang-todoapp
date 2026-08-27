package users_transport_http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/emp2ty0/golang-todoapp/internal/core/domain"
	core_logger "github.com/emp2ty0/golang-todoapp/internal/core/logger"
	core_http_request "github.com/emp2ty0/golang-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/emp2ty0/golang-todoapp/internal/core/transport/http/response"
	core_http_type "github.com/emp2ty0/golang-todoapp/internal/core/transport/http/types"
	core_http_utils "github.com/emp2ty0/golang-todoapp/internal/core/transport/http/utils"
)

type PathUserRequest struct {
	FullName    core_http_type.Nullable[string] `json:"full_name"`
	PhoneNumber core_http_type.Nullable[string] `json:"phone_number"`
}

func (r *PathUserRequest) Validate() error {
	if r.FullName.Set {
		if r.FullName.Value == nil {
			return fmt.Errorf("FullName cant be NULL ")
		}

		fullNameLen := len([]rune(*r.FullName.Value))

		if fullNameLen < 3 || fullNameLen > 100 {
			return fmt.Errorf("FullName must be beetween 3 and 100 symbols")
		}
	}

	if r.PhoneNumber.Set {
		if r.PhoneNumber.Value != nil {
			phoneNumberLen := len([]rune(*r.PhoneNumber.Value))

			if phoneNumberLen < 10 || phoneNumberLen > 15 {
				return fmt.Errorf("PhoneNumber must be beetwen 10 and 15 symbols")
			}

			if !strings.HasPrefix(*r.PhoneNumber.Value, "+") {
				return fmt.Errorf("PhoneNumber must be prefix +")
			}
		}
	}

	return nil
}

func (h *UserHTTPHandler) PathUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get userID path value")

		return
	}

	var request PathUserRequest
	if err := core_http_request.DecodeAndvalidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failde to decode and valudate HTTP request")

		return
	}

	userPatch := userPatchFromRequest(request)

	userDomain, err := h.userService.PatchUser(ctx, userID, userPatch)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to patch user")

		return
	}

	response := UserDTOFromDomain(userDomain)

	responseHandler.JSONResponse(response, http.StatusOK)

}

func userPatchFromRequest(request PathUserRequest) domain.UserPatch {
	return domain.UserPatch{
		FullName:    request.FullName.ToDomain(),
		PhoneNumber: request.PhoneNumber.ToDomain(),
	}
}
