package grpcsvc

import (
	"context"
	proto "multiply/grpc/protov1"

	"google.golang.org/grpc"
)

func NewClient(conn *grpc.ClientConn) *proto.MultiplyServiceClient {
	client := proto.NewMultiplyServiceClient(conn)
	return &client
}

func ExecuteOperation(x float32, y float32, client *proto.MultiplyServiceClient) (*float32, error) {
	ctx := context.Background()
	result, err := (*client).Multiply(ctx, &proto.Request{
		X: x,
		Y: y,
	})
	if err != nil {
		return nil, err
	}
	return &result.Z, nil
}
