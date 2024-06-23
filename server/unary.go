package main

import (
	"context"

	pb "github.com/Minz10/gRPC-go/proto"
)

func (s *helloServer) SayHello(ctx context.Context, re *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}