package client

import (
	"context"
	"fmt"

	"github.com/geothomas11/go-grpc-order-svc/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductServiceClient struct {
	Client pb.ProductServiceClient
}

func InitProductServiceClient(url string) ProductServiceClient {
	cc, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("could not connect:", err)
	}

	c := ProductServiceClient{
		Client: pb.NewProductServiceClient(cc),
	}
	return c
}
func (c *ProductServiceClient) FindOne(productId int64) (*pb.FindOneResponse, error) {
	req := &pb.FindOneRequest{
		Id: productId,
	}
	return c.Client.FindOne(context.Background(), req)

}

func (C *ProductServiceClient) DecreaseStock(productId int64, orderId int64) (*pb.DecreaseStockResponse, error) {
	req := &pb.DecreaseStockRequest{
		Id:      productId,
		OrderId: orderId,
	}
	return C.Client.DecreaseStock(context.Background(),req)

}
