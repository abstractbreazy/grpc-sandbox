package server

import (
	"context"

	protos "github.com/abstractbreazy/grpc-sandbox/payments/grpc/gen/proto/payments/v1"
	"github.com/hashicorp/go-hclog"
	_ "google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

//Compiler time check
var _ protos.PaymentsServer = (*Payments)(nil)

// Payments is a gRPC-server 
type Payments struct {
	log hclog.Logger
}

// NewPayments creates a new Payments server
func NewPayments(l hclog.Logger) *Payments {
	return &Payments{l}
}

// GetStatus implements the ExampleServer GetStatus method
func (c *Payments) Ping(ctx context.Context, _ *emptypb.Empty) (
	*protos.Pong, error) {
	c.log.Info("Ping!")		
	return &protos.Pong{Response: "Pong!"}, nil
}

func (c *Payments) Create(ctx context.Context, in *protos.Payment) (_ *emptypb.Empty, err error) {

	panic("implement me")
}

func (c *Payments) Delete(ctx context.Context, in *protos.DeletePayments) (_ *emptypb.Empty, err error) {

	panic("implement me")
}

func (c *Payments) List(ctx context.Context, _ *emptypb.Empty) (in *protos.PaymentsList, err error) {
	
	panic("implement me")
}