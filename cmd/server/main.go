package main

import (
	"fmt"
	"log"

	"github.com/go-chi/chi/v5"

	"github.com/latif-ecommerce-microservices/gateway-service/internal/config"
	"github.com/latif-ecommerce-microservices/gateway-service/internal/transport/grpc"
	"github.com/latif-ecommerce-microservices/gateway-service/internal/transport/grpc/client"
	httpTransport "github.com/latif-ecommerce-microservices/gateway-service/internal/transport/http"
	"github.com/latif-ecommerce-microservices/gateway-service/internal/transport/http/auth"
	"github.com/latif-ecommerce-microservices/gateway-service/internal/transport/http/user"
	"github.com/latif-ecommerce-microservices/gateway-service/pkg/logging"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	grpcConns, err := grpc.NewConnections(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer grpcConns.Close()

	logger := logging.CreateDefaultLogger(cfg.GetLogLevel().String())

	authClient := client.NewAuthClient(grpcConns.User)

	authHandler := auth.NewHandler(authClient, logger)
	userHandler := user.NewHandler()

	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {
		auth.RegisterRoutes(r, authHandler)
		user.RegisterRoutes(r, userHandler, cfg.JWTSecret)
	})

	server := httpTransport.NewServer(cfg, r)

	logger.Info(fmt.Sprintf(
		"[HTTP] API Gateway running on port %s...",
		cfg.AppHTTPPort,
	))

	log.Fatal(server.Start())
}
