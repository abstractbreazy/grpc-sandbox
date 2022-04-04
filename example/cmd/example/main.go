package main

import (
	"os"

	protos "github.com/abstractbreazy/grpc-sandbox/example/grpc/gen/proto/example/v1"
	examplesrv "github.com/abstractbreazy/grpc-sandbox/example/server"
	"github.com/abstractbreazy/grpc-sandbox/internal/server"

	"github.com/hashicorp/go-hclog"
)

func main() {
	var (
		log = hclog.Default()
		err error
	)

	// create an instance of the gRPC server.
	var srv *server.Server
	if srv, err = server.New(); err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	// create an instance of the Example server.
	cs := examplesrv.NewExample(log)

	// register the Example server.
	protos.RegisterExampleServer(srv.GRPC(), cs)

	// run gRPC server asynchronously.
	srv.RunAsync()
	defer srv.Close()

	log.Info("service starting on 9092...")

	server.Wait()

	log.Info("stopping server...")
}
