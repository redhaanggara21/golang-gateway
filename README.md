
Golang Gateway Service With Grpc

Base what i learn from this blog: 

https://levelup.gitconnected.com/microservices-with-go-grpc-api-gateway-and-authentication-part-1-2-393ad9fc9d30

thanks somuch and cheers.

making the day possible to do this architect for myown project.

check grpc:

1. compile one and one using command example: protoc -I pkg/product/pb/ pkg/product/pb/product.proto --go_out=plugins=grpc:.
    still not found how to generaly compiles
2. the another options: protoc --go_out=plugins=grpc:. **/*.proto
3. still find another way