package main

import (
	"log"

	"github.com/latif-ecommerce-microservices/gateway-service/internal/config"
	"github.com/latif-ecommerce-microservices/gateway-service/internal/router"
)

func main() {
	cfg := config.LoadConfig()

	app := router.Setup(cfg)

	log.Println("API Gateway running on port", cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}
