package main;

import (
	"context"
	"log"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	pb "github.com/saurabh10041998/grpc-go-course/calculator/proto"
)

func doSqrt(c pb.CalculatorServiceClient, number int32) {
	log.Printf("Invoking Sqrt function with a streaming request\n")
	
	req := &pb.SqrtRequest{
		Number: number,
	}

	res, err := c.Sqrt(context.Background(), req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			// user error
			log.Printf("Error message from server: %v\n", e.Message())
			log.Printf("Error code from server: %v\n", e.Code())
			if e.Code() == codes.InvalidArgument {
				log.Printf("We probably sent a negative number!\n")
			}
		} else {
			log.Fatalf("Big error calling Sqrt: %v\n", err)
		}
	} else {
		log.Printf("Result of Sqrt of %v: %v\n", number, res.GetResult())
	}
}