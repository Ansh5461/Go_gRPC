package main

import (
	"context"
	"log"

	pb "github.com/Ansh5461/Go_gRPC/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")

	//now we will call RPC endpoints
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Ansh ",
	})

	if err != nil {
		log.Fatalf("Could not greet %v\n", err)
	}

	log.Printf("Greeting : %s\n", res.Result)
}
