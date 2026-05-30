package main;

import (
	"context"
	"log"
	"time"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "github.com/saurabh10041998/grpc-go-course/greet/proto"
)

func doGreetWithDeadline(c pb.GreetServiceClient, deadline time.Duration) {
	log.Printf("Starting to do a GreetWithDeadline RPC with a deadline of %v\n", deadline)
	ctx, cancel := context.WithTimeout(context.Background(), deadline)
	defer cancel()
	
	req := &pb.GreetRequest{
		FirstName: "Saurabh",
	}

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Printf("Deadline was exceeded! Deadline was %v\n", deadline)
			} else {
				log.Printf("Unexpected error: %v\n", statusErr)
			}
		} else {
			log.Fatalf("Error while calling GreetWithDeadline: %v\n", err)
		}
		return
	}
	log.Printf("Response from GreetWithDeadline: %v\n", res.GetResult())
}