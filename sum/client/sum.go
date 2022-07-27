package main

import (
	"context"
	"log"

	pb "github.com/Ansh5461/Go_gRPC/sum/proto"
)

func dosum(c pb.SumServiceClient) {
	log.Println("Dosum function from client side invoked")

	res, err := c.Calculate(context.Background(), &pb.SumRequest{
		Num1: 25,
		Num2: 69,
	})

	if err != nil {
		log.Fatalf("Unable to sum %v\n", err)
	}

	//now client has called server, and server has sent the required information which is stored in res
	//print res data now

	log.Printf("Sum of 2 numbers is %v", res.Res)
}
