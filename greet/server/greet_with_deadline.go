package main;

import (
	"context"
	"log"
	"time"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/saurabh10041998/grpc-go-course/greet/proto"
)


func (s *Server) GreetWithDeadline(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("GreetWithDeadline function was invoked with %v\n", req)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			log.Printf("Client canceled the request\n")
			return nil, status.Error(codes.Canceled, "Client canceled the request")
		}
		time.Sleep(1 * time.Second)
	}
	
	firstName := req.GetFirstName()
	result := "Hello " + firstName
	return &pb.GreetResponse{
		Result: result,
	}, nil
}