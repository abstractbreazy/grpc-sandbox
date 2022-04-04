package main

import (
	"os"

	protos "github.com/abstractbreazy/grpc-sandbox/example/grpc/gen/proto/example/v1"
	examplesrv "github.com/abstractbreazy/grpc-sandbox/example/server"
	"github.com/abstractbreazy/grpc-sandbox/internal/server"

	"github.com/hashicorp/go-hclog"
	//"google.golang.org/grpc"
	//"google.golang.org/grpc/reflection"
)

// func main() {

// 	log := hclog.Default()

// 	// create a new gRPC server.
// 	gs := grpc.NewServer()

// 	// create an instance of the Example server.
// 	cs := server.NewExample(log)

// 	// register the Example server.
// 	protos.RegisterExampleServer(gs, cs)

// 	// register the reflection service which allows clients to deterimite the methods
// 	// for gRPC service.
// 	reflection.Register(gs)

// 	// create a TCP socket for incoming server connection.
// 	l, err := net.Listen("tcp", ":9092")
// 	if err != nil {
// 		log.Error("Unable to listen", "error", err)
// 		os.Exit(1)
// 	}
// 	log.Info("service starting on 9092...")

// 	gs.Serve(l)
// }


func main() {

	log := hclog.Default()
	var err error

	// create an instance of the gRPC server.
	var srv *server.Server
	if srv, err = server.New(); err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	// create an instance of the Example server.
    cs := examplesrv.NewExample(log)

	protos.RegisterExampleServer(srv.GRPC(), cs)

	// run gRPC server asynchronously.
	srv.RunAsync()
	defer srv.Close()

	log.Info("service starting on 9092...")

	server.Wait()

	log.Info("stopping server...")
}