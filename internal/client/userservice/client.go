// internal/client/user_service/client.go
package userservice

import (
	"context"
	authpb "github.com/latif-ecommerce-microservices/user-service/pkg/pb/auth"
	"google.golang.org/grpc"
)

type Client struct {
	auth authpb.AuthServiceClient
}

func NewClient(conn *grpc.ClientConn) *Client {
	return &Client{
		auth: authpb.NewAuthServiceClient(conn),
	}
}

func (c *Client) Login(ctx context.Context, email, password string) (*authpb.LoginResponse, error) {
	return c.auth.Login(ctx, &authpb.LoginRequest{
		Email:    email,
		Password: password,
	})
}
