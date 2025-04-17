package auth

import (
	"fmt"

	"github.com/geothomas11/go-grpc-api-gateway/pkg/auth/pb"
	"github.com/geothomas11/go-grpc-api-gateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.AuthServicePClient
}

func InitServicesClient(c *config.Config) pb.AuthServicePClient {
	cc, err := grpc.NewClient(
		c.ProductSvcUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		fmt.Println("coudnot connect:", err)
	}
	return pb.NewAuthServicePClient(cc)

}
