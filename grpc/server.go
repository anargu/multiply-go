package grpcsvc

import (
	"context"
	"multiply"
	"multiply/grpc/proto"
)

type grpcserver struct {
	proto.UnimplementedMultiplyServiceServer
}

func (s *grpcserver) Multiply(ctx context.Context, request *proto.Request) (*proto.Result, error) {
	x, y := request.GetX(), request.GetY()
	result, err := multiply.Multiply(&x, &y)
	if err != nil {
		return nil, err
	}

	return &proto.Result{Z: *result}, nil
}

var MultiplyGrpcSrv = &grpcserver{}
