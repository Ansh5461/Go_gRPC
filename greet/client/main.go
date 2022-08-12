package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Ansh5461/Go_gRPC/greet/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect %v", err)
	}

	defer conn.Close()

	//we will create an instance of greet service client
	c := pb.NewGreetServiceClient(conn)

	//doGreet will greet the client, only 1 times
	//doGreet(c)

	//this function will do server streaming
	//doGreetManyTimes(c)

	//this function is for client streaming
	//doLongGreet(c)

	//for client server streaming
	doGreetEveryone(c)
}
