package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/Ansh5461/Go_gRPC/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	//we are doing client streaming

	log.Println("LongGreet was invoked")

	//we will keep on receving stream from client, as client is streaming, then we will return the response, whatever we want to return

	res := ""
	//res is empty string but we will concatenate all the strings from client to it
	//this time client is sending the stream, so server needs to receive it
	for {
		req, err := stream.Recv()

		//what will happen when client won't have anything to send
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		log.Printf("Received the request as %v\n", req.FirstName)

		if err != nil {
			log.Fatalf("Got error while receiving stream %v\n", err)
		}

		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}

}
