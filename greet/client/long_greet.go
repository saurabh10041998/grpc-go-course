package main;

import (
	"log"
	"context"
	"time"
	pb "github.com/saurabh10041998/grpc-go-course/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Printf("Starting to do a Long Greet RPC...")
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v", err)
	}
	reqs := []*pb.GreetRequest{
		&pb.GreetRequest{
			FirstName: "Saurabh",
		},
		&pb.GreetRequest{
			FirstName: "John",
		},
		&pb.GreetRequest{
			FirstName: "Doe",
		},
	}
	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v", err)
	}
	log.Printf("LongGreet Response: %v\n", res)
}