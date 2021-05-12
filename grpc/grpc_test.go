package grpcsvc

import (
	"context"
	"log"
	"multiply/grpc/proto"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var (
	defaultPort = "7000"
	lis         *bufconn.Listener
)

func init() {
	lis = bufconn.Listen(bufSize)
	srv := grpc.NewServer()
	proto.RegisterMultiplyServiceServer(srv, MultiplyGrpcSrv)
	go func() {
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestGrpcFailedCall(t *testing.T) {
	differntPort := ":8000"
	conn, err := grpc.Dial(differntPort, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := NewClient(conn)
	var x float32 = 2
	var y float32 = 4
	z, err := ExecuteOperation(x, y, client)
	if err == nil {
		t.Fatalf("error shouldn't be nil: %v. response: %v\n", err, z)
	}
}
func TestGrpcCall(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := NewClient(conn)
	var x float32 = 2
	var y float32 = 4
	z, err := ExecuteOperation(x, y, client)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	log.Printf("Response: %+v", z)
}
