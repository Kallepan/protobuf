package service

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"time"

	greetpb "github.com/kallepan/protobuf/proto/greet/v1"
)

type GreetService struct {
	greetpb.UnimplementedGreetServiceServer
}

func (s *GreetService) HelloWorld(_ context.Context, req *greetpb.NoParam) (*greetpb.HelloResponse, error) {
	return &greetpb.HelloResponse{
		Message: "Hello",
	}, nil
}

func (s *GreetService) HelloWorldServerStreaming(req *greetpb.NamesList, stream greetpb.GreetService_HelloWorldServerStreamingServer) error {
	for _, name := range req.GetNames() {
		res := &greetpb.HelloResponse{
			Message: fmt.Sprintf("Hello %s", name),
		}

		if err := stream.Send(res); err != nil {
			return err
		}

		time.Sleep(2 * time.Second)
	}

	slog.Info("Finished streaming")

	return nil
}

func (s *GreetService) HelloWorldClientStreaming(stream greetpb.GreetService_HelloWorldClientStreamingServer) error {
	var names []string

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			slog.Info("Finished streaming")
			return stream.SendAndClose(&greetpb.MessagesList{Messages: names})
		}

		if err != nil {
			return err
		}

		slog.Info(fmt.Sprintf("Received name: %s", req.GetName()))
		names = append(names, req.GetName())
	}
}

func (s *GreetService) HelloWorldBidirectionalStreaming(stream greetpb.GreetService_HelloWorldBidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			slog.Info("Finished streaming")
			return nil
		}

		if err != nil {
			return err
		}

		slog.Info(fmt.Sprintf("Received name: %s", req.GetName()))

		res := &greetpb.HelloResponse{
			Message: fmt.Sprintf("Hello %s", req.GetName()),
		}

		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
