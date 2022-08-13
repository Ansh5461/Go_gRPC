package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Ansh5461/Go_gRPC/calculator/proto"
)

func doMaxFind(c pb.PrimeFindClient) {
	//here we will keep on sending numbers to server and keep on receiving numbers too, just with the use of 2 go routines
	log.Println("function doMaxFind was invoked")

	stream, err := c.MaxTillNow(context.Background())
	if err != nil {
		log.Fatalf("Got an error while calling the function: %v\n", err)
	}

	reqs := []*pb.Number{
		{Inp: 1},
		{Inp: 5},
		{Inp: 3},
		{Inp: 6},
		{Inp: 2},
		{Inp: 20},
	}

	waitc := make(chan struct{})

	//creating 2 goroutines each for sending and receiving requests

	//sending
	go func() {
		for _, req := range reqs {
			log.Printf("Sending number %d", req.Inp)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	//receiving
	go func() {
		//infinite receving till some condition is not met
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Got error while receiving %v", err)
				break
			}
			log.Printf("Received result : %v\n", res.Max)
			log.Println()
		}
		close(waitc)
	}()

	<-waitc
}
