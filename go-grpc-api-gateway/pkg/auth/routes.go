package auth

import (
	"github.com/geothomas11/go-grpc-api-gateway/pkg/auth/routes"
	"github.com/geothomas11/go-grpc-api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient{
	svc:=&ServiceClient{
		Client: InitServicesClient(c),
	}
	routes:=r.Group("/auth")
	routes.POST("/register",svc.)

}
