package main

import (
	"context"

	"log"

	pb "github.com/Ansh5461/Go_gRPC/greet/proto"
)

//this file will contain implementation of RPC endpoints
//finding method signature for grpc end points, go to generated code, and type greet round bracket open, and select a line under greet interface

//we want to implement this server under server file

//we have created the Greet function here, as we have defined in gRPC proto file, that this function will implement grpc end points
func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Function was invoked by %v\n", in)

	return &pb.GreetResponse{
		Result: "Hello my friend " + in.FirstName,
	}, nil
}
