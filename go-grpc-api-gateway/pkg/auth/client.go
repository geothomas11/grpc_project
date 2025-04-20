package auth

import (
	"log"

	"github.com/geothomas11/go-grpc-api-gateway/pkg/auth/pb"
	"github.com/geothomas11/go-grpc-api-gateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	conn, err := grpc.NewClient(c.AuthSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("couldn't connect to auth service: %v", err)
	}

	return pb.NewAuthServiceClient(conn)
}
