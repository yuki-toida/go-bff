package microservice_email

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"go-bff/email/pb"
	"google.golang.org/grpc"
)

type Service interface {
	Build(ctx context.Context, request *pb.BuildRequest) (*pb.BuildResponse, error)
	Reverse(ctx context.Context, request *pb.ReverseRequest) (*pb.ReverseResponse, error)
}

type service struct {
	build   endpoint.Endpoint
	reverse endpoint.Endpoint
}

func New(conn *grpc.ClientConn) Service {
	return &service{
		build: grpctransport.NewClient(
			conn,
			"proto.Email",
			"Build",
			func(_ context.Context, request interface{}) (interface{}, error) {
				return request.(*pb.BuildRequest), nil
			},
			func(_ context.Context, response interface{}) (interface{}, error) {
				return response.(*pb.BuildResponse), nil
			},
			pb.BuildResponse{},
		).Endpoint(),
		reverse: grpctransport.NewClient(
			conn,
			"proto.Email",
			"Reverse",
			func(_ context.Context, request interface{}) (interface{}, error) {
				return request.(*pb.ReverseRequest), nil
			},
			func(_ context.Context, response interface{}) (interface{}, error) {
				return response.(*pb.ReverseResponse), nil
			},
			pb.ReverseResponse{},
		).Endpoint(),
	}
}

func (s *service) Build(ctx context.Context, request *pb.BuildRequest) (*pb.BuildResponse, error) {
	res, err := s.build(ctx, request)
	if err != nil {
		return nil, err
	}
	return res.(*pb.BuildResponse), nil
}

func (s *service) Reverse(ctx context.Context, request *pb.ReverseRequest) (*pb.ReverseResponse, error) {
	res, err := s.reverse(ctx, request)
	if err != nil {
		return nil, err
	}
	return res.(*pb.ReverseResponse), nil
}
