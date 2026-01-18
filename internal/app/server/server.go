package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/latif-ecommerce-microservices/gateway-service/internal/config"
	"github.com/latif-ecommerce-microservices/gateway-service/internal/handler"
	"github.com/latif-ecommerce-microservices/gateway-service/pkg/logging"

	authpb "github.com/latif-ecommerce-microservices/user-service/pkg/pb/auth"
)

type Server struct {
	srv    *http.Server
	router chi.Router
	cfg    *config.Config
	logger *logging.Logger

	userConn *grpc.ClientConn

	authClient authpb.AuthServiceClient
}

func NewAppServer(cfg *config.Config, logger *logging.Logger) *Server {
	router := chi.NewMux()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.AppHTTPPort),
		Handler: router,
	}

	return &Server{
		cfg:    cfg,
		router: router,
		srv:    srv,
		logger: logger,
	}
}

func (s *Server) BeforeStart(ctx context.Context) error {
	conn, err := grpc.Dial(
		"localhost:8000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}

	s.userConn = conn
	s.authClient = authpb.NewAuthServiceClient(conn)

	authHandler := handler.NewAuthHandler(s.authClient, s.logger)

	s.router.Route("/api/v1", func(r chi.Router) {
		r.Post("/auth/login", authHandler.Login)
	})

	return nil
}

func (s *Server) AfterStart(ctx context.Context) error {
	if s.userConn != nil {
		return s.userConn.Close()
	}
	return nil
}

func (s *Server) HTTPServer() *http.Server {
	return s.srv
}
