package http

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/latif-ecommerce-microservices/gateway-service/internal/config"
	"net/http"
)

type Server struct {
	httpServer *http.Server
	router     chi.Router
}

func NewServer(
	cfg *config.Config,
	router chi.Router,
) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    ":" + cfg.AppHTTPPort,
			Handler: router,
		},
		router: router,
	}
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
