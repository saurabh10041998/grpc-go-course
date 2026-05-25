package main;

import (
	"log"
	pb "github.com/saurabh10041998/grpc-go-course/primes/proto"
)

func (s *Server) GetPrimes(req *pb.PrimeRequest, stream pb.PrimeService_GetPrimesServer) error {
	log.Printf("Received request for number: %v\n", req.Number);
	var k int32 = 2;
	n := req.Number;

	for n > 1 {
		if n % k == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: k,
			})
			n = n / k;
		} else {
			k++;
		}
	}
	return nil;
}