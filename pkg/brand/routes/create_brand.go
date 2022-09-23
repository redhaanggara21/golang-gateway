package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redhaanggara21/go-grpc-api-gateway/pkg/brand/pb"
)

type CreateBrandRequestBody struct {
	brandName   string `json:"name"`
	dateCreated string `json:"date"`
}

func CreateBrand(ctx *gin.Context, c pb.BrandServiceClient) {
	b := CreateBrandRequestBody{}

	b.brandName = "a"
	b.dateCreated = "120x"

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateBrand(
		context.Background(),
		&pb.CreateBrandRequest{
			BrandName:   b.brandName,
			DateCreated: b.dateCreated,
		},
	)

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
