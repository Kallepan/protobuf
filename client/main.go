package main

import (
	greetpb "github.com/kallepan/protobuf/proto/greet/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	con, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer con.Close()

	client := greetpb.NewGreetServiceClient(con)

	// call the service
	names := &greetpb.NamesList{
		Names: []string{"Kalle", "Pelle", "Nisse"},
	}

	callHelloWorld(client)
	callHelloWorldServerStream(client, names)
	callHelloWorldClientStream(client, names)
	callHelloWorldBiDiStream(client, names)
}
