package main

import (
	"fmt"
	"log"
	"net"

	"github.com/geothomas11/go-grpc-auth-svc/pkg/config"
	"github.com/geothomas11/go-grpc-auth-svc/pkg/db"
	"github.com/geothomas11/go-grpc-auth-svc/pkg/pb"
	"github.com/geothomas11/go-grpc-auth-svc/pkg/services"
	"github.com/geothomas11/go-grpc-auth-svc/pkg/utils"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("failed at config", err)
	}
	dbUrl := c.GetDBUrl()

	h := db.Init(dbUrl)

	jwt := utils.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: 24 * 365,
	}
	lis, err := net.Listen("tcp", ":"+c.Port)

	if err != nil {
		log.Fatalln("Falied to listing:", c.Port)
	}
	fmt.Println("Auth svc is on", c.Port)
	s := services.Server{
		H:   h,
		JWT: jwt,
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to sereve:", err)
	}

}
