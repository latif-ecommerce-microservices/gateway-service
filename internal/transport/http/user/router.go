package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/latif-ecommerce-microservices/gateway-service/internal/middleware"
)

func RegisterRoutes(
	r chi.Router,
	h *Handler,
	jwtSecret string,
) {
	r.Route("/users", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(jwtSecret))
		r.Get("/me", h.Me)
	})
}
