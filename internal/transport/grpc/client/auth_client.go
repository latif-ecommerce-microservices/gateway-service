package client

import (
	"context"

	authpb "github.com/latif-ecommerce-microservices/user-service/pkg/pb/auth"
	"google.golang.org/grpc"
)

type AuthClient interface {
	Login(ctx context.Context, email, password string) (*authpb.LoginResponse, error)
}

type authClient struct {
	client authpb.AuthServiceClient
}

func NewAuthClient(conn *grpc.ClientConn) AuthClient {
	return &authClient{
		client: authpb.NewAuthServiceClient(conn),
	}
}

func (a *authClient) Login(
	ctx context.Context,
	email string,
	password string,
) (*authpb.LoginResponse, error) {
	return a.client.Login(ctx, &authpb.LoginRequest{
		Email:    email,
		Password: password,
	})
}
