package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Ansh5461/Go_gRPC/greet/proto"
)

//we will create a function which will call server to request for some information, which in our case is, printing hello many times
//it will call rpc end points many times, from there it will receive a rpc stream and error, if error then log, less
//infinite loop calling stream.Receive, if any EOF then break, else keep printing the loop
func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("goGreetManyTimes invoked")

	//create a greetrequest instance
	req := &pb.GreetRequest{
		FirstName: " Ansh",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling greet many times %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Got the error while reading from the stream %v\n", err)
		}

		log.Printf("GreetManyTimes  %s\n", msg.Result)
	}
}
