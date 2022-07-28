package main

import (
	pb "github.com/Ansh5461/Go_gRPC/calculator/proto"

	"log"
)

func (s *Server) Primes(in *pb.PrimeRequest, stream pb.PrimeFind_PrimesServer) error {

	log.Printf("Primes function was invoked by %v\n", in)

	/*Pseudo code for this is

	k = 2
	N = 210
	while N > 1:
		if N % k == 0:   // if k evenly divides into N
			print k      // this is a factor
			N = N / k    // divide N by k so that we have the rest of the number left.
		else:
			k = k + 1
	*/

	k := int64(2)
	num := in.Num

	/*This function was written by me, other function was written by instructor, I went on with his solution because its neet
	for{
		if num % k == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: k,
			})

			num = num / k
		} else {
			k = k + 1
		}

		if num <= 1{
			break
		}
	}*/

	for num > 1 {
		if num%k == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: k,
			})

			num = num / k
		} else {
			k = k + 1
		}
	}

	return nil
}
