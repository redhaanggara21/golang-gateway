package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redhaanggara21/go-grpc-api-gateway/pkg/product/pb"
)

type CreateProductRequestBody struct {
	Name  string `json:"name"`
	SKU   string `json:"sku"`
	Stock int64  `json:"stock"`
	Price int64  `json:"price"`
}

func CreateProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	b := CreateProductRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateProduct(
		context.Background(),
		&pb.CreateProductRequest{
			Name:  b.Name,
			Sku:   b.SKU,
			Stock: b.Stock,
			Price: b.Price,
		},
	)

	fmt.Print(&res)

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
