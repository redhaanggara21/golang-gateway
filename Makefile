proto:
    protoc --go_out=plugins=grpc:. **/*.proto
    
server:
	go run cmd/main.go