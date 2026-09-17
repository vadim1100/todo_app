package core_http_server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	core_logger "github.com/vadim1100/todo_app/internal/core/logger"
	core_http_middleware "github.com/vadim1100/todo_app/internal/core/transport/http/middleware"
	"go.uber.org/zap"
)

type HTTPServer struct {
	mux *http.ServeMux
	config Config
	logger *core_logger.Logger
	middleware []core_http_middleware.Middleware
}

func NewHTTPServer(
	config Config,
	logger *core_logger.Logger,
	middleware ...core_http_middleware.Middleware,
) *HTTPServer {
	return &HTTPServer{
		mux: http.NewServeMux(),
		config: config,
		logger: logger,
		middleware: middleware,
	}
}

func (h *HTTPServer) Run(ctx context.Context) error{
	mux := core_http_middleware.ChainMiddleware(h.mux, h.middleware...)

	server := &http.Server{
		Addr: h.config.Addr,
		Handler: mux,
	}

	ch := make(chan error, 1)

	go func ()  {
		defer close(ch)

		h.logger.Warn("start http server", zap.String("addr", h.config.Addr))

		err := server.ListenAndServe()

		if !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("listen and serve http: %w", err)
		}
	case <-ctx.Done():
		h.logger.Warn("shutdown http server...")

		shutdownContext, cancel := context.WithTimeout(context.Background(), h.config.ShutdownTimeout)

		defer cancel()

		if err := server.Shutdown(shutdownContext); err != nil {
			_ = server.Close()
			return fmt.Errorf("shutdown http server: %w", err)
		}

		h.logger.Warn("http server stopped")
	}
	return nil
}

func (s *HTTPServer) RegisterRouters(routers ...http.Handler) {
    for _, router := range routers {
		s.mux.Handle("/", router)
    }
}