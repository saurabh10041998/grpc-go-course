package main;

import (
	"net"
	"log"
	"google.golang.org/grpc"
	pb "github.com/saurabh10041998/grpc-go-course/greet/proto"
)

var addr = "0.0.0.0:50051";

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s\n", err)
	}
}