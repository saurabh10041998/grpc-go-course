package main;

import (
	"context"
	"io"
	"log"
	"time"
	pb "github.com/saurabh10041998/grpc-go-course/calculator/proto"
)


func doMax(c pb.CalculatorServiceClient) {
	log.Printf("Starting to do a Max Bidi Streaming RPC...\n")
	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
		return
	}
	
	waitc := make(chan struct{})
	// send go routine
	go func() {
		numbers := []int32{1,5,3,6,2,20}
		for _, number := range numbers {
			log.Printf("Sending number: %v\n", number)
			stream.Send(&pb.MaxRequest{
				Number: number,
			})
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	// receive go routine
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				log.Printf("No more data from server\n")
				break
			}
			if err != nil {
				log.Printf("Error while receiving data: %v\n", err)
				break
			}
			log.Printf("Received new max value: %v\n", res.GetMax())
		}
		close(waitc)
	}()
	
	<-waitc

}