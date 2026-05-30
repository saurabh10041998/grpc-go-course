package main;

import (
	"io"
	"log"
	pb "github.com/saurabh10041998/grpc-go-course/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone function was invoked with a streaming request")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("Error while reading client stream: %v\n", err)
			return err
		}
		log.Printf("Received request: %v\n", req)
		res := &pb.GreetResponse{
			Result: "Hello " + req.GetFirstName(),
		}
		if err := stream.Send(res); err != nil {
			log.Printf("Error while sending response to client: %v\n", err)
			return err
		}
	}
}