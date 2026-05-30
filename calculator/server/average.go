package main;

import (
	"io"
	"log"
	pb "github.com/saurabh10041998/grpc-go-course/calculator/proto"
)

func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {
	log.Printf("Average function was invoked\n");
	var sum float32 = 0
	var count int32 = 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			average := sum / float32(count)
			return stream.SendAndClose(&pb.AverageResponse{
				Average: average,
			})
		}
		if err != nil {
			log.Printf("Error while reading client stream: %v\n", err)
			return err
		}
		log.Printf("Received number: %v\n", req.GetNumber())
		sum += req.GetNumber()
		count++
	}
}