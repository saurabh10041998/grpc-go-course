package main;

import (
	"context"
	"log"
	"time"
	pb "github.com/saurabh10041998/grpc-go-course/calculator/proto"
)

func doAverage(c pb.CalculatorServiceClient) {
	log.Printf("Starting to do a streaming RPC for Average...\n")
	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Average: %v\n", err)
	}
	numbers := []float32{3, 5, 9, 54, 23}
	for _, number := range numbers {
		log.Printf("Sending number: %v\n", number)
		stream.Send(&pb.AverageRequest{
			Number: number,
		})
		time.Sleep(1 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from Average: %v\n", err)
	}
	log.Printf("The average is: %v\n", res.GetAverage())
}