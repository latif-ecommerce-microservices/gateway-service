package handler

import (
	"encoding/json"
	"net/http"

	"github.com/latif-ecommerce-microservices/gateway-service/internal/transport/grpc/client"
	"github.com/latif-ecommerce-microservices/gateway-service/pkg/httputil"
	"github.com/latif-ecommerce-microservices/gateway-service/pkg/logging"
)

type AuthHandler struct {
	authClient client.AuthClient
	logger     *logging.Logger
}

func NewAuthHandler(
	authClient client.AuthClient,
	logger *logging.Logger,
) *AuthHandler {
	return &AuthHandler{
		authClient: authClient,
		logger:     logger,
	}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httputil.WriteBadRequestResponse(w, "Invalid request body", nil)
		return
	}

	res, err := h.authClient.Login(
		r.Context(),
		req.Email,
		req.Password,
	)
	if err != nil {
		httputil.HandleError(w, h.logger, err)
		return
	}

	data := map[string]string{
		"access_token":  res.AccessToken,
		"refresh_token": res.RefreshToken,
	}

	httputil.WriteSuccessResponse(w, data, "Login successful")
}
