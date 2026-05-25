package main;

import (
	"log"
	"context"
	pb "github.com/saurabh10041998/grpc-go-course/calculator/proto"
)

func doCalculate(c pb.CalculatorServiceClient) {
	log.Printf("Calculating 3 + 5\n");
	res, err := c.Add(context.Background(), &pb.CalculateRequest{
		A: 3,
		B: 5,
	});
	
	if err != nil {
		log.Fatalf("Error while calling Add RPC: %v\n", err)
	}
	log.Printf("Response from Add: %v\n", res.Result);
}