// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ExampleClient is the client API for Example service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExampleClient interface {
	// Ping returns the <-pong message.
	Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PingResponse, error)
}

type exampleClient struct {
	cc grpc.ClientConnInterface
}

func NewExampleClient(cc grpc.ClientConnInterface) ExampleClient {
	return &exampleClient{cc}
}

func (c *exampleClient) Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/example.v1.Example/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleServer is the server API for Example service.
// All implementations should embed UnimplementedExampleServer
// for forward compatibility
type ExampleServer interface {
	// Ping returns the <-pong message.
	Ping(context.Context, *emptypb.Empty) (*PingResponse, error)
}

// UnimplementedExampleServer should be embedded to have forward compatible implementations.
type UnimplementedExampleServer struct {
}

func (UnimplementedExampleServer) Ping(context.Context, *emptypb.Empty) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

// UnsafeExampleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExampleServer will
// result in compilation errors.
type UnsafeExampleServer interface {
	mustEmbedUnimplementedExampleServer()
}

func RegisterExampleServer(s grpc.ServiceRegistrar, srv ExampleServer) {
	s.RegisterService(&Example_ServiceDesc, srv)
}

func _Example_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.v1.Example/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServer).Ping(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Example_ServiceDesc is the grpc.ServiceDesc for Example service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Example_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "example.v1.Example",
	HandlerType: (*ExampleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Example_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/v1/example.proto",
}
