package microservice_email

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"go-bff/email/pb"
	"google.golang.org/grpc"
)

type Service interface {
	Build(ctx context.Context, request *pb_email.BuildRequest) (*pb_email.BuildResponse, error)
	Reverse(ctx context.Context, request *pb_email.ReverseRequest) (*pb_email.ReverseResponse, error)
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
				return request.(*pb_email.BuildRequest), nil
			},
			func(_ context.Context, response interface{}) (interface{}, error) {
				return response.(*pb_email.BuildResponse), nil
			},
			pb_email.BuildResponse{},
		).Endpoint(),
		reverse: grpctransport.NewClient(
			conn,
			"proto.Email",
			"Reverse",
			func(_ context.Context, request interface{}) (interface{}, error) {
				return request.(*pb_email.ReverseRequest), nil
			},
			func(_ context.Context, response interface{}) (interface{}, error) {
				return response.(*pb_email.ReverseResponse), nil
			},
			pb_email.ReverseResponse{},
		).Endpoint(),
	}
}

func (s *service) Build(ctx context.Context, request *pb_email.BuildRequest) (*pb_email.BuildResponse, error) {
	res, err := s.build(ctx, request)
	if err != nil {
		return nil, err
	}
	return res.(*pb_email.BuildResponse), nil
}

func (s *service) Reverse(ctx context.Context, request *pb_email.ReverseRequest) (*pb_email.ReverseResponse, error) {
	res, err := s.reverse(ctx, request)
	if err != nil {
		return nil, err
	}
	return res.(*pb_email.ReverseResponse), nil
}
