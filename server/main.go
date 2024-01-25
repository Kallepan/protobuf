package main

import (
	"log/slog"
	"net"

	greetpb "github.com/kallepan/protobuf/proto/greet/v1"
	"github.com/kallepan/protobuf/server/config"
	"github.com/kallepan/protobuf/server/service"
	"google.golang.org/grpc"
)

func main() {
	config.InitLogger()

	// open port
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		slog.Error(err.Error())
	}

	// launch the server
	grpcServer := grpc.NewServer()

	// register the service
	greetpb.RegisterGreetServiceServer(grpcServer, &service.GreetService{})

	// launch the server
	if err := grpcServer.Serve(listener); err != nil {
		slog.Error(err.Error())
	}
}
