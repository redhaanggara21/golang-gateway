proto:
	PATH=$PATH:$GOPATH/bin/ protoc pkg/**/pb/*.proto --go_out=plugins=grpc:.

# protoc \ -I=../protos \ ../protos/**/*.proto \ pkg/**/pb/*.proto

# PATH=$PATH:$GOPATH/bin/ protoc pkg/**/pb/*.proto --go_out=plugins=grpc:.

# PATH=$PATH:$GOPATH/bin/ protoc --go_out=. *.proto
# PATH=$PATH:$GOPATH/bin/ protoc --proto_path=pkg/product/pb pkg/product/pb/*.proto --go_out=plugins=grpc:.

# use this for one and one from outside dir
#protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pkg/product/pb/*.proto

#protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pkg/**/pb/*.proto

# PATH=$PATH:$GOPATH/bin/ protoc --go_out=. *.proto

# protoc --go_out=. --go_opt=paths=source_relative \ --go-grpc_out=. --go-grpc_opt=paths=source_relative \ pkg/product/pb/*.proto
	
server:
	go run cmd/main.go

# gen:
#     protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb

# clean:
#     rm pb/*.go 

# run:
#     go run main.go

# protoc pkg/**/pb/*.proto --go_out=plugins=grpc:.