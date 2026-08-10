package core_http_server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	core_logger "github.com/emp2ty0/golang-todoapp/internal/core/logger"
	core_http_middleware "github.com/emp2ty0/golang-todoapp/internal/core/transport/http/middleware"
	"go.uber.org/zap"
)

type HTTPServer struct {
	mux *http.ServeMux

	config Config
	log    *core_logger.Logger

	midleware []core_http_middleware.Middleware
}

func NewHTTPServer(config Config, log *core_logger.Logger, midleware ...core_http_middleware.Middleware) *HTTPServer {
	return &HTTPServer{
		mux:       http.NewServeMux(),
		config:    config,
		log:       log,
		midleware: midleware,
	}
}

func (h *HTTPServer) RegisterAPIRouters(routers ...*APIVersionRouter) {
	for _, router := range routers {
		prefix := "/api/" + string(router.apiVersion)

		h.mux.Handle(prefix+"/", http.StripPrefix(prefix, router))
	}
}

func (h *HTTPServer) Run(ctx context.Context) error {
	mux := core_http_middleware.ChainMidleware(h.mux, h.midleware...)

	server := &http.Server{
		Addr:    h.config.Addr,
		Handler: mux,
	}

	ch := make(chan error, 1)

	go func() {
		defer close(ch)

		err := server.ListenAndServe()

		if !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	h.log.Warn("Start HTTP server", zap.String("addr", h.config.Addr))

	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("listen and serve http %w", err)
		}
	case <-ctx.Done():
		h.log.Warn("shutdown HTTP server")

		shutDownCtx, cancel := context.WithTimeout(context.Background(), h.config.ShutdownTimeount)

		defer cancel()

		if err := server.Shutdown(shutDownCtx); err != nil {
			_ = server.Close()

			return fmt.Errorf("shutdown HTTP server %w", err)

		}

		h.log.Warn("HTTP server stopped")
	}

	return nil

}
