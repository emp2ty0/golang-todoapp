package core_http_response

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func (h *HTTPResponseHandler) PanicResponse(p any, msg string) {
	statusCode := http.StatusInternalServerError
	err := fmt.Errorf("unexpected panuc %v", p)

	h.log.Error(msg, zap.Error(err))
	h.rw.WriteHeader(statusCode)

	reponse := map[string]string{
		"messange": msg,
		"error":    err.Error(),
	}

	if err := json.NewEncoder(h.rw).Encode(reponse); err != nil {
		h.log.Error("write HTTP error", zap.Error(err))
	}

}
