package main;

import (
	"context"
	"io"
	"log"
	"time"
	pb "github.com/saurabh10041998/grpc-go-course/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("Starting to do a BiDi streaming RPC...")
	
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
		return
	}
	waitc := make(chan struct{})
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				// we've reached the end of the stream
				break
			}
			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}
			log.Printf("Received: %v\n", res.GetResult())
		}
		close(waitc)
	}()
	requests := []*pb.GreetRequest{
		{FirstName: "Saurabh"},
		{FirstName: "John"},
		{FirstName: "Jane"},
		{FirstName: "Doe"},
	}
	for _, req := range requests {
		log.Printf("Sending: %v\n", req)
		if err := stream.Send(req); err != nil {
			log.Printf("Error while sending: %v\n", err)
			return
		}
		time.Sleep(1 * time.Second)
	}
	stream.CloseSend()
	<-waitc
}