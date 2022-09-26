package brand

import (
	"fmt"

	"github.com/redhaanggara21/go-grpc-api-gateway/pkg/brand/pb"
	"github.com/redhaanggara21/go-grpc-api-gateway/pkg/config"
	"google.golang.org/grpc"
)

type BrandClient struct {
	Client pb.BrandServiceClient
}

func InitBrandClient(c *config.Config) pb.BrandServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.BrandSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewBrandServiceClient(cc)
}
