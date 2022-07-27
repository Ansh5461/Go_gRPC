package main

import (
	"log"
	"net"

	pb "github.com/Ansh5461/Go_gRPC/sum/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.SumServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Error while connecting %v", err)
	}

	log.Printf("Listening on %v", addr)

	s := grpc.NewServer()

	pb.RegisterSumServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v\n", err)
	}
}

//A gRPC service definition uses protocol buffers as its Interface Definition Language to describe the service
