package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Ansh5461/Go_gRPC/sum/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Unable to connect %v", err)
	}

	defer conn.Close()

	//c - client, of the given type
	c := pb.NewSumServiceClient(conn)

	dosum(c)
}
