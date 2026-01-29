package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/latif-ecommerce-microservices/gateway-service/internal/transport/grpc/client"
	"github.com/latif-ecommerce-microservices/gateway-service/pkg/logging"
	"net/http"

	"github.com/latif-ecommerce-microservices/gateway-service/pkg/httputil"
)

type Handler struct {
	userClient client.UserClient
	logger     *logging.Logger
}

func NewHandler(
	userClient client.UserClient,
	logger *logging.Logger,
) *Handler {
	return &Handler{
		userClient: userClient,
		logger:     logger,
	}
}

func (h *Handler) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		httputil.WriteBadRequestResponse(w, "user id is required", nil)
		return
	}

	res, err := h.userClient.GetUserById(
		r.Context(),
		id,
	)
	if err != nil {
		httputil.HandleError(w, h.logger, err)
		return
	}

	data := map[string]string{
		"id":    res.Id,
		"name":  res.Name,
		"email": res.Email,
	}

	httputil.WriteSuccessResponse(w, data, "Successfully retrieved user")
}
