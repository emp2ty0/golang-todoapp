package core_http_response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	core_errors "github.com/emp2ty0/golang-todoapp/internal/core/errors"
	core_logger "github.com/emp2ty0/golang-todoapp/internal/core/logger"
	"go.uber.org/zap"
)

type HTTPResponseHandler struct {
	log *core_logger.Logger
	rw  http.ResponseWriter
}

func NewHTTPResponseHandler(log *core_logger.Logger, rw http.ResponseWriter) *HTTPResponseHandler {
	return &HTTPResponseHandler{
		log: log,
		rw:  rw,
	}
}

func (h *HTTPResponseHandler) JSONResponse(
	responseBody any,
	statusCode int,
) {
	h.rw.WriteHeader(statusCode)

	if err := json.NewEncoder(h.rw).Encode(responseBody); err != nil {
		h.log.Error("write HTTP response", zap.Error(err))
	}
}

func (h *HTTPResponseHandler) ErrorResponse(err error, msg string) {
	var (
		statucCode int
		logFunc    func(string, ...zap.Field)
	)

	switch {
	case errors.Is(err, core_errors.ErrInvalidArgument):
		statucCode = http.StatusBadGateway
		logFunc = h.log.Warn

	case errors.Is(err, core_errors.ErrNotFound):
		statucCode = http.StatusNotFound
		logFunc = h.log.Debug

	case errors.Is(err, core_errors.ErrConflict):
		statucCode = http.StatusConflict
		logFunc = h.log.Warn

	default:
		statucCode = http.StatusInternalServerError
		logFunc = h.log.Error
	}

	logFunc(msg, zap.Error(err))

	h.errorResponse(statucCode, err, msg)
}

func (h *HTTPResponseHandler) PanicResponse(p any, msg string) {
	statusCode := http.StatusInternalServerError
	err := fmt.Errorf("unexpected panuc %v", p)

	h.log.Error(msg, zap.Error(err))

	h.errorResponse(statusCode, err, msg)

}

func (h *HTTPResponseHandler) errorResponse(
	statusCode int,
	err error,
	msg string,
) {
	h.rw.WriteHeader(statusCode)

	reponse := map[string]string{
		"messange": msg,
		"error":    err.Error(),
	}

	h.JSONResponse(reponse, statusCode)

}
