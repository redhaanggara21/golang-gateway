package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/redhaanggara21/go-grpc-api-gateway/pkg/auth"
	"github.com/redhaanggara21/go-grpc-api-gateway/pkg/brand"
	"github.com/redhaanggara21/go-grpc-api-gateway/pkg/config"
	"github.com/redhaanggara21/go-grpc-api-gateway/pkg/order"
	"github.com/redhaanggara21/go-grpc-api-gateway/pkg/product"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	fmt.Print(&c)

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)
	brand.RegisterRoutes(r, &c, &authSvc)
	//some command and check in

	r.Run(c.Port)
}
