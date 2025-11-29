package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/latif-ecommerce-microservices/gateway-service/internal/config"
	"github.com/latif-ecommerce-microservices/gateway-service/internal/middleware"
	"github.com/latif-ecommerce-microservices/gateway-service/internal/proxy"
)

func Setup(cfg *config.Config) *fiber.App {
	app := fiber.New()

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// Apply JWT middleware
	app.Use(middleware.Auth())

	// User service
	app.All("/users/*", proxy.Forward(cfg.UserService))

	// Order service
	app.All("/orders/*", proxy.Forward(cfg.OrderService))

	return app
}
