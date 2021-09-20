package main

import (
	"net"
	"os"

	protos "github.com/AbstractBreazy/grpc-sandbox/example/grpc/gen/proto/go/example/v1"
	"github.com/AbstractBreazy/grpc-sandbox/example/server"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	// create a new gRPC server
	gs := grpc.NewServer()

	// create an instance of the Example server
	cs := server.NewExample(log)

	// register the Example server 
	protos.RegisterExampleServer(gs, cs)

	// register the reflection service which allows clients to deterimite the methods 
	// for gRPC service.
	reflection.Register(gs)

	// create a TCP socket for incoming server connection.
	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}
	log.Info("service starting...")
	gs.Serve(l)
}