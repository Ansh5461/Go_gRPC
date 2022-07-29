package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Ansh5461/Go_gRPC/calculator/proto"
)

func doAverage(c pb.PrimeFindClient) {
	log.Printf("doAverage was invoked\n")

	//I am back and I definitely wont allow anything to come in between
	nums := []*pb.Inp{
		{Num: 142},
		{Num: 731},
		{Num: 786},
		{Num: 635},
		{Num: 458},
	}

	stream, err := c.Average(context.Background())

	if err != nil {
		log.Fatalf("Got error while calling average function%v\n", err)
	}

	for _, num := range nums {
		log.Printf("Sending %v to the server \n", num)
		stream.Send(num)
		time.Sleep(1 * time.Second)
	}

	//now we have sent all the values to server via stream
	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Printf("Got error while receiving the answer %v\n", err)
	}

	log.Printf("Average of numbers is %v\n", res.Res)
}
