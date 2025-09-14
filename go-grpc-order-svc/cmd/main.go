package main

import (
	"log"
	"net"

	"github.com/geothomas11/go-grpc-order-svc/pkg/client"
	"github.com/geothomas11/go-grpc-order-svc/pkg/config"
	"github.com/geothomas11/go-grpc-order-svc/pkg/db"
	"github.com/geothomas11/go-grpc-order-svc/pkg/pb"
	"github.com/geothomas11/go-grpc-order-svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
		log.SetFlags(log.LstdFlags | log.Lshortfile) // Adds filename and line number
	}

	h := db.Init(c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	productSvc := client.InitProductServiceClient(c.ProductSvcUrl)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	log.Println("Order Svc on", c.Port)

	s := services.Server{
		H:          h,
		ProductSvc: productSvc,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterOrderServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
