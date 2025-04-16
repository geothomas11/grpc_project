package product

import (
	"github.com/geothomas11/go-grpc-api-gateway/pkg/auth"
	"github.com/geothomas11/go-grpc-api-gateway/pkg/auth/routes"
	"github.com/geothomas11/go-grpc-api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)
	svc := &auth.ServiceClient{
		Client: auth.InitServicesClient(c),
	}

	routes := r.Group("/product")
	routes.Use(a.AuthRequired)
	routes.POST("/createproduct", svc.CreateProduct)
	routes.GET("/:id", svc.FindOne)
}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FindOne(ctx, svc.Client)
}

func (svc*ServiceClient)CreateProduct(ctx*gin.Context)  {
	routes.CreateProduct(ctx,svc.Client)
	
}
