package main

import (
	"fmt"
	"log"

	pb "github.com/Ansh5461/Go_gRPC/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked by parameters : %v\n", in)

	//we will implement a loop 10 times and we will use send keyword
	//send keyword is defined in stream variable of type pb.GreetService_GreetManyTimesServer
	for i := 0; i < 10; i++ {
		//res is variable we will create which is formatted version which we will put in result

		res := fmt.Sprintf("Hello %s times %v", in.FirstName, i)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}
	return nil
}

//go build -o bin/greet/server.exe ./greet/server
