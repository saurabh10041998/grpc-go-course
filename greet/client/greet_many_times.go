package main;

import (
	"io"
	"log"
	"context"
	pb "github.com/saurabh10041998/grpc-go-course/greet/proto"
)


func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Printf("Starting to do a streaming RPC...\n")
	req	 := &pb.GreetRequest{
		FirstName: "Saurabh",
	}
	resStream, err := c.GreetManyTimes(context.Background(), req);
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes RPC: %v\n", err)
	}
	for {
		msg, err := resStream.Recv();
		if err == io.EOF {
			// we've reached the end of the stream
			break;
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err)
		}
		log.Printf("Response from GreetManyTimes: %v\n", msg.GetResult());
	}
}