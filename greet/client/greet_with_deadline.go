package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Ansh5461/Go_gRPC/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	//this function will be called if the deadline is exceed
	defer cancel()

	//now context has enough information and now we can call our RPC endpoints
	req := &pb.GreetRequest{
		FirstName: "Ansh",
	}

	//here instead of calling with context.Background, we will be calling with our new context that we have created
	res, err := c.GreetWithDeadline(ctx, req)

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			//it means this is a gRPC error
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded")
				return
			} else {
				log.Fatalf("Unexpected error: %v\n", err)
			}
		} else {
			log.Fatalf("A non gRPC error %v\n", err)
		}
	}

	log.Printf("Result : %s\n", res.Result)
}
