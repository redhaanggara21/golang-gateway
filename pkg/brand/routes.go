package brand

import (
	"github.com/gin-gonic/gin"
	"github.com/redhaanggara21/go-grpc-api-gateway/pkg/auth"
	"github.com/redhaanggara21/go-grpc-api-gateway/pkg/brand/routes"
	"github.com/redhaanggara21/go-grpc-api-gateway/pkg/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &BrandClient{
		Client: InitBrandClient(c),
	}

	routes := r.Group("/brand")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateBrand)
	routes.GET("/:id", svc.FindOne)
}

func (svc *BrandClient) FindOne(ctx *gin.Context) {
	routes.FindOne(ctx, svc.Client)
}

func (svc *BrandClient) CreateBrand(ctx *gin.Context) {
	routes.CreateBrand(ctx, svc.Client)
}
