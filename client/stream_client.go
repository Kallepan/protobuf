package main

import (
	"context"
	"fmt"
	"log/slog"

	greetpb "github.com/kallepan/protobuf/proto/greet/v1"
)

func callHelloWorldClientStream(client greetpb.GreetServiceClient, names *greetpb.NamesList) {
	slog.Info("Starting to do a Client Streaming RPC...")

	stream, err := client.HelloWorldClientStreaming(context.Background())
	if err != nil {
		panic(err)
	}

	for _, name := range names.GetNames() {
		req := &greetpb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			panic(err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		panic(err)
	}

	for _, msg := range res.GetMessages() {
		slog.Info(fmt.Sprintf("Response: %s", msg))
	}
}
