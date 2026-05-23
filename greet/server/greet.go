package main;
import (
	"log"
	"context"
	pb "github.com/saurabh10041998/grpc-go-course/greet/proto"
)

func (s Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in);
	return &pb.GreetResponse {
		Result: "Hello " + in.FirstName,
	}, nil
}