package main

import (
	"context"
	"io"
	"log/slog"

	greetpb "github.com/kallepan/protobuf/proto/greet/v1"
)

func callHelloWorldBiDiStream(client greetpb.GreetServiceClient, names *greetpb.NamesList) {
	slog.Info("Starting to do a Client Streaming RPC...")

	stream, err := client.HelloWorldBidirectionalStreaming(context.Background())
	if err != nil {
		panic(err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}

			slog.Info(message.GetMessage())

			if err != nil {
				slog.Error(err.Error())
				break
			}
		}
		close(waitc)
	}()

	for _, name := range names.GetNames() {
		req := &greetpb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			panic(err)
		}
	}
	stream.CloseSend()
	<-waitc
	slog.Info("Finished streaming")
}
