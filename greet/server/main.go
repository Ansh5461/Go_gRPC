package main

import (
	"log"
	"net"

	pb "github.com/Ansh5461/Go_gRPC/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

//protoc -Igreet/proto --go_out=. --go_opt=module=github.com/Ansh5461/Go_gRPC --go-grpc_out=. --go-grpc_opt=module=github.com/Ansh5461/Go_gRPC greet/proto/dummy.proto

type Server struct {
	//here we will refer to that server that was created by gRPC server compiler
	//we are going to use this server to implement all the gRPC end points we will create in our greet proto
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	//before we use this listener, we are going to create a gRPC server object
	//done doing that

	s := grpc.NewServer()

	//here we will register greet service server
	//we will pass server to this function and pass an instance of type server struct to it

	//grpc server s needs an instance of greet service server, and we will use Server struct to implement
	//all grpc end points
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v\n", err)
	}

}
