package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/geothomas11/go-grpc-api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type RegisterRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	email := ctx.Query("email")
	password := ctx.Query("password")

	if email == "" || password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		fmt.Println("err", err)
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(int(res.Status), &res)
}
