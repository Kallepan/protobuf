package main

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	greetpb "github.com/kallepan/protobuf/proto/greet/v1"
)

func callHelloWorld(client greetpb.GreetServiceClient) {
	slog.Info("Starting to do a Unary RPC...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// call the service
	res, err := client.HelloWorld(ctx, &greetpb.NoParam{})
	if err != nil {
		panic(err)
	}

	// print the response
	slog.Info(fmt.Sprintf("Response: %s", res.GetMessage()))
}
