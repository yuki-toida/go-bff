package main

import (
	"context"
	"github.com/go-kit/kit/log"
	"go-bff/email/pb"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	// The gRPC listener mounts the Go kit gRPC server we created.
	grpcListener, err := net.Listen("tcp", ":8081")
	if err != nil {
		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
		os.Exit(1)
	}
	defer grpcListener.Close()

	server := grpc.NewServer()
	pb_email.RegisterEmailServer(server, &Server{})
	logger.Log("err", server.Serve(grpcListener))
}

type Server struct{}

func (s *Server) Build(_ context.Context, r *pb_email.BuildRequest) (*pb_email.BuildResponse, error) {
	address := r.Email + "@hacobu.jp"
	return &pb_email.BuildResponse{EmailAddress: address}, nil
}

func (s *Server) Reverse(_ context.Context, r *pb_email.ReverseRequest) (*pb_email.ReverseResponse, error) {
	runes := []rune(r.Email)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return &pb_email.ReverseResponse{EmailAddress: string(runes)}, nil
}
