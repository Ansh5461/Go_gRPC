package main

import (
	"context"
	"log"

	pb "github.com/Ansh5461/Go_gRPC/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.PrimeFindClient, n int32) {
	log.Println("function doSqrt was invoked")

	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{Number: n})

	if err != nil {
		e, ok := status.FromError(err)

		//ok will tell if there is a gRPC error
		if ok {
			//now e contans 2 things, message and code
			log.Printf("Error message from server %s\n", e.Message())
			log.Printf("Error code from server %s\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Println("We probably sent a negative number! ")
				return
			}
		} else {
			log.Fatalf("A non gRPC error %v\n", err)
		}
	}

	log.Printf("Square root of a number is %v\n", res.Result)
}
