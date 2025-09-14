package order

import (
	"log"

	"github.com/geothomas11/go-grpc-api-gateway/pkg/config"
	"github.com/geothomas11/go-grpc-api-gateway/pkg/order/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.OrderServiceClient
}

func InitServiceClient(c *config.Config) pb.OrderServiceClient {

	cc, err := grpc.NewClient(c.OrderSvcUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("Could not connect:", err)
	}

	return pb.NewOrderServiceClient(cc)
}
