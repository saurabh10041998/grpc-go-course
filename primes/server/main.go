package main;

import (
	"log"
	"net"
	"google.golang.org/grpc"
	pb "github.com/saurabh10041998/grpc-go-course/primes/proto"
)

var addr string = "localhost:50051";

type Server struct {
	pb.PrimeServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr);
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	s := grpc.NewServer()
	pb.RegisterPrimeServiceServer(s, &Server{});
	log.Printf("Server is listening on %v\n", addr);
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}	
}