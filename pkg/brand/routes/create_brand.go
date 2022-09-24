package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redhaanggara21/go-grpc-api-gateway/pkg/brand/pb"
)

type CreateBrandRequestBody struct {
	BrandName   string `json:"brandName"`
	DateCreated string `json:"dateCreated"`
}

func CreateBrand(ctx *gin.Context, c pb.BrandServiceClient) {
	b := CreateBrandRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateBrand(
		context.Background(),
		&pb.CreateBrandRequest{
			BrandName:   b.BrandName,
			DateCreated: b.DateCreated,
		},
	)

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
