package server_test

import (
	"context"
	"os"
	"testing"

	protos "github.com/abstractbreazy/grpc-sandbox/example/grpc/gen/proto/example/v1"
	srv "github.com/abstractbreazy/grpc-sandbox/pkg/testutils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func dialAddress() (address string) {
	const defaultDialAddress = "127.0.0.1:15010"
	if address = os.Getenv("TEST_DIAL_ADDRESS"); address != "" {
		return
	}

	return defaultDialAddress
}

type client struct {
	conn    *grpc.ClientConn
	example protos.ExampleClient
}

func newClient(t testing.TB, address string) (c *client) {

	c = new(client)
	var err error

	c.conn, err = grpc.Dial(address, grpc.WithInsecure())
	require.NoError(t, err)

	c.example = protos.NewExampleClient(c.conn)
	return
}

func (c *client) Close() (err error) {
	return c.conn.Close()
}

func TestGRPC_Ping(t *testing.T) {

	var (
		addr  = dialAddress()
		grpcs srv.GRPCServer
	)
	require.NoError(t, grpcs.Start(addr))
	defer grpcs.Close()

	var c = newClient(t, addr)
	defer c.Close()

	var (
		ctx         = context.Background()
		err         error
		pingMessage *protos.PingResponse
	)

	pingMessage, err = c.example.Ping(ctx, &emptypb.Empty{})
	assert.NoError(t, err)

	assert.EqualValues(t,
		protos.PingResponse{Message: "Pong!"}.Message,
		pingMessage.Message)
}
