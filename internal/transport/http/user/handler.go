package user

import (
	"net/http"

	"github.com/latif-ecommerce-microservices/gateway-service/pkg/httputil"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	// contoh: ambil user dari context (JWT)
	// user := middleware.GetUser(r.Context())

	httputil.WriteSuccessResponse(w, map[string]string{
		"message": "user profile",
	}, "OK")
}
