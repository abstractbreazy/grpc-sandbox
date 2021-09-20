package server

import (
	"context"
    protos "github.com/AbstractBreazy/grpc-sandbox/example/grpc/gen/proto/go/example/v1"
	"github.com/hashicorp/go-hclog"
	_ "google.golang.org/grpc"
    "google.golang.org/protobuf/types/known/emptypb"
)

// Compiler time check
var _ protos.ExampleServer = (*Example)(nil)

// Example is a gRPC-server 
type Example struct {
	log hclog.Logger
}

// NewExample creates a new Example server
func NewExample(l hclog.Logger) *Example {
	return &Example{l}
}

// GetStatus implements the ExampleServer GetStatus method
func (c *Example) GetStatus(ctx context.Context, _ *emptypb.Empty) (
	*protos.StatusResponse, error) {
	c.log.Info("example.v1.Example.GetStatus method")		
	return &protos.StatusResponse{Status: "OK"}, nil
}