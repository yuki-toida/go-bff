package grpc

import (
	"fmt"
	"google.golang.org/grpc"
)

func Connect(host, port string) *grpc.ClientConn {
	// gRPC client
	target := fmt.Sprintf("%s:%s", host, port)
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return conn
}
