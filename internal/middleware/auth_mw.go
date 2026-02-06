package middleware

import (
	"context"
	"fmt"
	"github.com/latif-ecommerce-microservices/gateway-service/pkg/httputil"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const (
	UserIDKey contextKey = "user_id"
)

func AuthMiddleware(jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				httputil.WriteUnauthorizedResponse(w, "unauthorized")
				return
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

			claims := jwt.MapClaims{}
			token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method")
				}
				return []byte(jwtSecret), nil
			})

			if err != nil || !token.Valid {
				httputil.WriteUnauthorizedResponse(w, "invalid token")
				return
			}

			userID, ok := claims["user_id"].(string)
			if !ok {
				httputil.WriteUnauthorizedResponse(w, "user_id missing")
				return
			}

			ctx := context.WithValue(r.Context(), UserIDKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
