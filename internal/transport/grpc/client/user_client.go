package client

import (
	"context"
	userpb "github.com/latif-ecommerce-microservices/user-service/pkg/pb/user"
	"google.golang.org/grpc"
)

type UserClient interface {
	GetUserById(ctx context.Context, id string) (*userpb.UserResponse, error)
}

type userClient struct {
	client userpb.UserServiceClient
}

func NewUserClient(conn *grpc.ClientConn) UserClient {
	return &userClient{
		client: userpb.NewUserServiceClient(conn),
	}
}

// List Endpoint
// ambil semua user (admin)
// ambil detail user (admin/user tbs)

func (a *userClient) GetUserById(
	ctx context.Context,
	id string,
) (*userpb.UserResponse, error) {
	return a.client.GetUserByID(ctx, &userpb.GetUserByIDRequest{
		Id: id,
	})
}
