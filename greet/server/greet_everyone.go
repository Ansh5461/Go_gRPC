package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/Ansh5461/Go_gRPC/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("Function GreetEveryOne was invoked")

	//we will do a combination of server streaming and client streaming
	//infite loop receiving something from the client, check EOF, if we receive something from client, we will directly send it back

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Got an error while receiving request, reading client stream %v", err)
		}

		res := fmt.Sprintln("Hello ", req.FirstName, " !")

		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Got error while sending it back %v", err)
		}
	}
}
