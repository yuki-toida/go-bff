package entity_notify

import (
	"context"
	"go-bff/notify/pb"
)

type Service interface {
	Build(ctx context.Context, request *pb.BuildRequest) (*pb.BuildResponse, error)
	Reverse(ctx context.Context, request *pb.ReverseRequest) (*pb.ReverseResponse, error)
}
