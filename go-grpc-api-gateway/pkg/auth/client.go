package auth

import (
	"fmt"

	"github.com/geothomas11/go-grpc-api-gateway/pkg/auth/pb"
	"github.com/geothomas11/go-grpc-api-gateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServicesClient(c *config.Config) pb.AuthServiceClient {
	cc, err := grpc.NewClient(
		c.AuthSvcUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		fmt.Println("coudnot connect:", err)
	}
	return pb.NewAuthServiceClient(cc)

}
