// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/greet/v1/greet.proto

package greetpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetServiceClient interface {
	HelloWorld(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*HelloResponse, error)
	HelloWorldServerStreaming(ctx context.Context, in *NamesList, opts ...grpc.CallOption) (GreetService_HelloWorldServerStreamingClient, error)
	HelloWorldClientStreaming(ctx context.Context, opts ...grpc.CallOption) (GreetService_HelloWorldClientStreamingClient, error)
	HelloWorldBidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (GreetService_HelloWorldBidirectionalStreamingClient, error)
}

type greetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetServiceClient(cc grpc.ClientConnInterface) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) HelloWorld(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/proto.greet.v1.GreetService/HelloWorld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetServiceClient) HelloWorldServerStreaming(ctx context.Context, in *NamesList, opts ...grpc.CallOption) (GreetService_HelloWorldServerStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[0], "/proto.greet.v1.GreetService/HelloWorldServerStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceHelloWorldServerStreamingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetService_HelloWorldServerStreamingClient interface {
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type greetServiceHelloWorldServerStreamingClient struct {
	grpc.ClientStream
}

func (x *greetServiceHelloWorldServerStreamingClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) HelloWorldClientStreaming(ctx context.Context, opts ...grpc.CallOption) (GreetService_HelloWorldClientStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[1], "/proto.greet.v1.GreetService/HelloWorldClientStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceHelloWorldClientStreamingClient{stream}
	return x, nil
}

type GreetService_HelloWorldClientStreamingClient interface {
	Send(*HelloRequest) error
	CloseAndRecv() (*MessagesList, error)
	grpc.ClientStream
}

type greetServiceHelloWorldClientStreamingClient struct {
	grpc.ClientStream
}

func (x *greetServiceHelloWorldClientStreamingClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceHelloWorldClientStreamingClient) CloseAndRecv() (*MessagesList, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(MessagesList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) HelloWorldBidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (GreetService_HelloWorldBidirectionalStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[2], "/proto.greet.v1.GreetService/HelloWorldBidirectionalStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceHelloWorldBidirectionalStreamingClient{stream}
	return x, nil
}

type GreetService_HelloWorldBidirectionalStreamingClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type greetServiceHelloWorldBidirectionalStreamingClient struct {
	grpc.ClientStream
}

func (x *greetServiceHelloWorldBidirectionalStreamingClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceHelloWorldBidirectionalStreamingClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetServiceServer is the server API for GreetService service.
// All implementations must embed UnimplementedGreetServiceServer
// for forward compatibility
type GreetServiceServer interface {
	HelloWorld(context.Context, *NoParam) (*HelloResponse, error)
	HelloWorldServerStreaming(*NamesList, GreetService_HelloWorldServerStreamingServer) error
	HelloWorldClientStreaming(GreetService_HelloWorldClientStreamingServer) error
	HelloWorldBidirectionalStreaming(GreetService_HelloWorldBidirectionalStreamingServer) error
	mustEmbedUnimplementedGreetServiceServer()
}

// UnimplementedGreetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (UnimplementedGreetServiceServer) HelloWorld(context.Context, *NoParam) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloWorld not implemented")
}
func (UnimplementedGreetServiceServer) HelloWorldServerStreaming(*NamesList, GreetService_HelloWorldServerStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method HelloWorldServerStreaming not implemented")
}
func (UnimplementedGreetServiceServer) HelloWorldClientStreaming(GreetService_HelloWorldClientStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method HelloWorldClientStreaming not implemented")
}
func (UnimplementedGreetServiceServer) HelloWorldBidirectionalStreaming(GreetService_HelloWorldBidirectionalStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method HelloWorldBidirectionalStreaming not implemented")
}
func (UnimplementedGreetServiceServer) mustEmbedUnimplementedGreetServiceServer() {}

// UnsafeGreetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetServiceServer will
// result in compilation errors.
type UnsafeGreetServiceServer interface {
	mustEmbedUnimplementedGreetServiceServer()
}

func RegisterGreetServiceServer(s grpc.ServiceRegistrar, srv GreetServiceServer) {
	s.RegisterService(&GreetService_ServiceDesc, srv)
}

func _GreetService_HelloWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).HelloWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.greet.v1.GreetService/HelloWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).HelloWorld(ctx, req.(*NoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreetService_HelloWorldServerStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NamesList)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServiceServer).HelloWorldServerStreaming(m, &greetServiceHelloWorldServerStreamingServer{stream})
}

type GreetService_HelloWorldServerStreamingServer interface {
	Send(*HelloResponse) error
	grpc.ServerStream
}

type greetServiceHelloWorldServerStreamingServer struct {
	grpc.ServerStream
}

func (x *greetServiceHelloWorldServerStreamingServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GreetService_HelloWorldClientStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).HelloWorldClientStreaming(&greetServiceHelloWorldClientStreamingServer{stream})
}

type GreetService_HelloWorldClientStreamingServer interface {
	SendAndClose(*MessagesList) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greetServiceHelloWorldClientStreamingServer struct {
	grpc.ServerStream
}

func (x *greetServiceHelloWorldClientStreamingServer) SendAndClose(m *MessagesList) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceHelloWorldClientStreamingServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GreetService_HelloWorldBidirectionalStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).HelloWorldBidirectionalStreaming(&greetServiceHelloWorldBidirectionalStreamingServer{stream})
}

type GreetService_HelloWorldBidirectionalStreamingServer interface {
	Send(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greetServiceHelloWorldBidirectionalStreamingServer struct {
	grpc.ServerStream
}

func (x *greetServiceHelloWorldBidirectionalStreamingServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceHelloWorldBidirectionalStreamingServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetService_ServiceDesc is the grpc.ServiceDesc for GreetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.greet.v1.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloWorld",
			Handler:    _GreetService_HelloWorld_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "HelloWorldServerStreaming",
			Handler:       _GreetService_HelloWorldServerStreaming_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "HelloWorldClientStreaming",
			Handler:       _GreetService_HelloWorldClientStreaming_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "HelloWorldBidirectionalStreaming",
			Handler:       _GreetService_HelloWorldBidirectionalStreaming_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/greet/v1/greet.proto",
}
