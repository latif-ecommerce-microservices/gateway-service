package handler

import (
	"encoding/json"
	"net/http"

	"github.com/latif-ecommerce-microservices/gateway-service/pkg/httputil"
	"github.com/latif-ecommerce-microservices/gateway-service/pkg/logging"

	authpb "github.com/latif-ecommerce-microservices/user-service/pkg/pb/auth"
)

type AuthHandler struct {
	grpcClient authpb.AuthServiceClient
	logger     *logging.Logger
}

func NewAuthHandler(grpcClient authpb.AuthServiceClient, logger *logging.Logger) *AuthHandler {
	return &AuthHandler{
		grpcClient: grpcClient,
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

	grpcRes, err := h.grpcClient.Login(r.Context(), &authpb.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		httputil.HandleError(w, h.logger, err)
		return
	}

	data := map[string]string{
		"access_token":  grpcRes.AccessToken,
		"refresh_token": grpcRes.RefreshToken,
	}

	httputil.WriteSuccessResponse(w, data, "Login successful")
}
