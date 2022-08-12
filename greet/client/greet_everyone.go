package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Ansh5461/Go_gRPC/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	//
	log.Println("Func doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream:  %v", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Ansh"},
		{FirstName: "Sharanya"},
		{FirstName: "Tarun"},
	}

	//now we will create a go channel, and we will have 2 goroutines will be running simultaneously

	//first goroutine - send all requests we have in this array
	//second goroutine - receive all the responses from the server

	waitc := make(chan struct{})

	go func() {
		//in this goroutine, we will have a loop that iterates over reqs and sends a request
		for _, req := range reqs {
			log.Printf("Sending request : %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
		//now we have sent all our requests to server
	}()

	go func() {
		//we will keep on receiving till server keeps on sending

		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Got error while receiving stream %v\n", err)
				break
			}

			log.Printf("Received : %v\n", res.Result)
		}
		close(waitc)
	}()

	//it will wait till these 2 go routines send and receive answers
	<-waitc

}
