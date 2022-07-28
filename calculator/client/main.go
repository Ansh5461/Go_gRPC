package main

import (
	pb "github.com/Ansh5461/Go_gRPC/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"log"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect %v\n", err)
	}

	defer conn.Close()

	c := pb.NewPrimeFindClient(conn)

	//call the function from here, by passing the parameter as c
	doPrimes(c)
}
