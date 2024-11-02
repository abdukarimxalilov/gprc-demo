package main

import (
	"context"
	pb "github.com/abdukarimxalilov/grpc-demo/proto"
)

func (s *helloServer) sayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error){
	return &pb.HelloResponse{
		Message: "Hello", 
	}, nil
}