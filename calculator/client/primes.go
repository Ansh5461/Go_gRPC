package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Ansh5461/Go_gRPC/calculator/proto"
)

func doPrimes(c pb.PrimeFindClient) {
	log.Printf("Function doPrimes was invoked")

	req := &pb.PrimeRequest{
		Num: 8478541021400,
	}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("Got error while sending the number %v\n", err)
	}

	//now if no error then it means we will receive a stream of answers, lets write for loop and wait for it
	log.Println("Primes of the digit are ")
	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Got error while receiving the data %v\n", err)
		}

		log.Printf("%d\t", res.Result)
	}
}
