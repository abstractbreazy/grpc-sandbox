package server_test

import (
	"context"
	"net"
	"testing"

	protos "github.com/abstractbreazy/grpc-sandbox/example/grpc/gen/proto/example/v1"
	examplesrv "github.com/abstractbreazy/grpc-sandbox/example/server"

	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/types/known/emptypb"
)

var listener *bufconn.Listener

func init() {
	listener = bufconn.Listen(1024 * 1024)
	s := grpc.NewServer()

	exampleServer := examplesrv.NewExample(hclog.NewNullLogger())
	protos.RegisterExampleServer(s, exampleServer)

	go func() {
		if err := s.Serve(listener); err != nil {
			panic(err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return listener.Dial()
}

func TestGRPC_Ping(t *testing.T) {
	t.Run("success ping", func(t *testing.T) {
		conn, err := grpc.DialContext(
			context.Background(),
			"bufnet",
			grpc.WithContextDialer(bufDialer),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		require.NoError(t, err)
		defer conn.Close()

		client := protos.NewExampleClient(conn)

		resp, err := client.Ping(context.Background(), &emptypb.Empty{})
		assert.NoError(t, err)
		assert.Equal(t, "Pong!", resp.Message)
	})
}
