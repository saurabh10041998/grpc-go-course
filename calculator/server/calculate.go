package main;

import (
	"log"
	"context"
	pb "github.com/saurabh10041998/grpc-go-course/calculator/proto"
)

func (s Server) Add(ctx context.Context, in *pb.CalculateRequest) (*pb.CalculateResponse, error) {
	log.Printf("Add is called with %v", in);
	return &pb.CalculateResponse{
		Result: in.A + in.B,
	}, nil
}