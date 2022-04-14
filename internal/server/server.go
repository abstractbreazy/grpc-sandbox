package server

import (
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server initialized and wraps a gRPC server.
type Server struct {
	l   net.Listener
	srv *grpc.Server
	wg  sync.WaitGroup
}

// New initiates a new gRPC instance with listener.
func New() (s *Server, err error) {
	s = new(Server)

	// create a TCP socket for incoming server connection.
	if s.l, err = net.Listen("tcp", ":9092"); err != nil {
		return
	}
	// create a new gRPC server.
	s.srv = grpc.NewServer()

	return
}

// Run an gRPC server in a goroutine asynchronously.
func (s *Server) RunAsync() {
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

// Stop the server gracefully.
func (s *Server) Close() (err error) {
	s.srv.GracefulStop()
	err = s.l.Close()
	s.wg.Wait()
	return
}

// GRPC server instance. Use this method to register a service implementation.
func (s *Server) GRPC() *grpc.Server {
	return s.srv
}

// waiting os signals.
func Wait() {
	var sig = make(chan os.Signal, 2)
	signal.Notify(sig, os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM)
	<-sig
}
