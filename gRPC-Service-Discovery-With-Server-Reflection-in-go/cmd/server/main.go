//go:generate protoc -I ../../proto --go_out=plugins=grpc:../../proto ../../proto/goheroe.proto
package main

import (
	"fmt"
	"net"

	"github.com/chemidy/goheroe-code-examples/gRPC-Service-Discovery-With-Server-Reflection-in-go/app"
	"github.com/chemidy/goheroe-code-examples/gRPC-Service-Discovery-With-Server-Reflection-in-go/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPort = ":44444"
)

func main() {

	// Creates a tcp listener on port :authGrpcPort
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
	}

	// Creates a new gRPC server
	grpcServer := grpc.NewServer()

	// get instance of GoHeroeServer
	srv, err := app.NewGoHeroeServer()
	if err != nil {
		fmt.Printf("failed to create GoHeroeServer : %s\n", err)
	}

	// register GoHeroeServer on grpc server
	superpower.RegisterGoHeroeServer(grpcServer, srv)

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	// start server
	fmt.Printf("starting server on port : %s\n", grpcPort)
	grpcServer.Serve(lis)
}
