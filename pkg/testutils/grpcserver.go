package tests

import (
	"net"
	"sync"

	protos "github.com/abstractbreazy/grpc-sandbox/example/grpc/gen/proto/example/v1"
	example_server "github.com/abstractbreazy/grpc-sandbox/example/server"
	"github.com/hashicorp/go-hclog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// a gRPC-Server for tests.
type GRPCServer struct {
	l   net.Listener
	srv *grpc.Server
	wg  sync.WaitGroup
}

func (grpcs *GRPCServer) Start(address string) (err error) {

	var l net.Listener
	if l, err = net.Listen("tcp", address); err != nil {
		return
	}

	var srv = grpc.NewServer()

	grpcs.l = l
	grpcs.srv = srv

	log := hclog.Default()

	cs := example_server.NewExample(log)

	// register the Example server.
	protos.RegisterExampleServer(srv, cs)

	grpcs.RunAsync()

	return
}

func (s *GRPCServer) Close() (err error) {
	s.srv.GracefulStop()
	err = s.l.Close()
	s.wg.Wait()
	return
}

func (s *GRPCServer) RunAsync() {
	// register the reflection service which allows clients to deterimite the methods
	// for gRPC service.
	reflection.Register(s.srv)
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		if err := s.srv.Serve(s.l); err != nil && err != grpc.ErrServerStopped {
			return
		}
	}()
}
