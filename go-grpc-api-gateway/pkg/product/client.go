package product

import (
	"log"

	"github.com/geothomas11/go-grpc-api-gateway/pkg/config"
	"github.com/geothomas11/go-grpc-api-gateway/pkg/product/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.ProductServiceClient
}

func InitServiceClient(c *config.Config) pb.ProductServiceClient {

	cc, err := grpc.NewClient(
		c.ProductSvcUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println("Could not connect:", err)
	}

	return pb.NewProductServiceClient(cc)
}
