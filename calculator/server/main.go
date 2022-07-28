package main

import (
	"log"
	"net"

	pb "github.com/Ansh5461/Go_gRPC/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

//whatever our serice will be there, we will implement the server of that service
type Server struct {
	pb.PrimeFindServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Got error while connecting  %v\n", err)
	}

	log.Printf("Listening on address %v\n", addr)

	//create a new grpc server

	s := grpc.NewServer()

	//register it as type server struct, means s has now the data of server struct which is our calculator type
	pb.RegisterPrimeFindServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v\n", err)
	}
}
