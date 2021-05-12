package grpcsvc

import (
	"context"
	"multiply"
	protov1 "multiply/grpc/protov1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcserverv1 struct {
	protov1.UnimplementedMultiplyServiceServer
}

func (s *grpcserverv1) Multiply(ctx context.Context, request *protov1.Request) (*protov1.Result, error) {
	x, y := request.GetX(), request.GetY()
	result, err := multiply.Multiply(&x, &y)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &protov1.Result{Z: *result}, nil
}

var MultiplyGrpcSrvV1 = &grpcserverv1{}

func RegisterGrpcServer(srv *grpc.Server) {
	protov1.RegisterMultiplyServiceServer(srv, MultiplyGrpcSrvV1)
	// able to add move proto service versions
}
