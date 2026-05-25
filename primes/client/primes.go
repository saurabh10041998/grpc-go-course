package main;

import (
	"io"
	"log"
	"context"
	pb "github.com/saurabh10041998/grpc-go-course/primes/proto"
)

func getPrimes(c pb.PrimeServiceClient) {
	log.Println("getPrimes was invoked");

	resStream, err := c.GetPrimes(context.Background(), &pb.PrimeRequest{
		Number: 120,
	})
	if err != nil {
		log.Fatalf("Error while calling GetPrimes: %v", err)
	}
	log.Println("Response from GetPrimes:");
	for {
		resp, err := resStream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break;
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		log.Printf("Prime factor: %v", resp.GetResult())
	}

}	