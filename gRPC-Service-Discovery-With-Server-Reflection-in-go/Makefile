all:
	protoc -I ./proto/ ./proto/goheroe.proto --go_out=plugins=grpc:proto

server:
	go run cmd/server/main.go