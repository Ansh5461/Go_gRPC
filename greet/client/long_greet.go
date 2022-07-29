package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Ansh5461/Go_gRPC/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Ansh"},
		{FirstName: "Sharanya"},
		{FirstName: "Tarun"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Got the error while calling function %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Got error while receiving answer %v\n", err)
	}

	log.Printf("LongGreet %s\n", res.Result)
}
