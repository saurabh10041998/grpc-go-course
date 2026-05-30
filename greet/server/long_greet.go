package main;

import (
	log "log"
	"fmt"
	"io"
	pb "github.com/saurabh10041998/grpc-go-course/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Printf("LongGreet function was invoked");
	res := ""
	for {
		req, err := stream.Recv();
		
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		firstName := req.FirstName;
		res += fmt.Sprintf("Hello %s!\n", firstName);
	}
}