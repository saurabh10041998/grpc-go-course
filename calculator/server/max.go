package main;

import (
	"log"
	"io"
	pb "github.com/saurabh10041998/grpc-go-course/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Printf("Max function was invoked with a streaming request\n")

	var max int32 = 0
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			log.Printf("Client has finished sending data\n")
			return nil
		}

		if err != nil {
			log.Printf("Error while reading client stream: %v\n", err)
			return err
		}

		number := req.GetNumber()
		log.Printf("Received number: %v\n", number)

		if number > max {
			max = number
			sendErr := stream.Send(&pb.MaxResponse{
				Max: max,
			})
			if sendErr != nil {
				log.Printf("Error while sending data to client: %v\n", sendErr)
				return sendErr
			}
		}
	}
}