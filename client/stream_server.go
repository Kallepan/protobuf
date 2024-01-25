package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"time"

	greetpb "github.com/kallepan/protobuf/proto/greet/v1"
)

func callHelloWorldServerStream(client greetpb.GreetServiceClient, names *greetpb.NamesList) {
	slog.Info("Starting to do a Server Streaming RPC...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.HelloWorldServerStreaming(ctx, names)
	if err != nil {
		panic(err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		slog.Info(fmt.Sprintf("Response: %s", res.GetMessage()))
	}
}
