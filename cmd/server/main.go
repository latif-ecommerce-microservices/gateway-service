package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/latif-ecommerce-microservices/gateway-service/internal/app/server"
	"github.com/latif-ecommerce-microservices/gateway-service/internal/config"
	"github.com/latif-ecommerce-microservices/gateway-service/pkg/logging"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	logger := logging.CreateDefaultLogger(cfg.GetLogLevel().String())

	appServer := server.NewAppServer(cfg, logger)

	if err := appServer.BeforeStart(ctx); err != nil {
		logger.Fatal(fmt.Sprintf("Failed to init gateway: %s", err.Error()))
	}

	go func() {
		logger.Info(fmt.Sprintf(
			"[HTTP] API Gateway running on port %s...",
			cfg.AppHTTPPort,
		))

		if err := appServer.HTTPServer().ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			logger.Fatal(fmt.Sprintf("HTTP server error: %v", err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down API Gateway...")

	if err := appServer.HTTPServer().Shutdown(ctx); err != nil {
		logger.Error(fmt.Sprintf("HTTP shutdown error: %v", err))
	}

	if err := appServer.AfterStart(ctx); err != nil {
		logger.Error(fmt.Sprintf("Cleanup error: %v", err))
	}

	logger.Info("API Gateway stopped gracefully")
}
