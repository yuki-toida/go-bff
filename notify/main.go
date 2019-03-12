package main

import (
	"context"
	"github.com/go-kit/kit/log"
	"go-bff/notify/pb"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	// The gRPC listener mounts the Go kit gRPC server we created.
	grpcListener, err := net.Listen("tcp", ":8081")
	if err != nil {
		os.Exit(1)
	}
	defer grpcListener.Close()

	server := grpc.NewServer()
	pb.RegisterEmailServer(server, &Server{})
	if err := logger.Log("err", server.Serve(grpcListener)); err != nil {
		panic(err)
	}
}

type Server struct{}

func (s *Server) Build(_ context.Context, r *pb.BuildRequest) (*pb.BuildResponse, error) {
	address := r.Email + "@hacobu.jp"
	return &pb.BuildResponse{EmailAddress: address}, nil
}

func (s *Server) Reverse(_ context.Context, r *pb.ReverseRequest) (*pb.ReverseResponse, error) {
	runes := []rune(r.Email)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return &pb.ReverseResponse{EmailAddress: string(runes)}, nil
}
