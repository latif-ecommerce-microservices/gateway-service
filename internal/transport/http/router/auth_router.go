package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/latif-ecommerce-microservices/gateway-service/internal/transport/http/handler"
)

func RegisterAuthRoutes(r chi.Router, h *handler.AuthHandler) {
	r.Post("/auth/login", h.Login)
}
