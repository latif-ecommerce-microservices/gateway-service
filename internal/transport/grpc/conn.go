package grpc

import (
	"context"
	"net"
	"time"

	"github.com/latif-ecommerce-microservices/gateway-service/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Connections struct {
	User *grpc.ClientConn
}

func NewConnections(cfg *config.Config) (*Connections, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	addr := net.JoinHostPort(
		cfg.AppHost,
		cfg.UserGRPCPort,
	)

	conn, err := grpc.DialContext(
		ctx,
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &Connections{
		User: conn,
	}, nil
}

func (c *Connections) Close() error {
	if c.User != nil {
		return c.User.Close()
	}
	return nil
}
