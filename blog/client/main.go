package main

import (
	pb "github.com/Ansh5461/Go_gRPC/blog/proto"
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

	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)

	readBlog(c, id)                   //this is valid
	readBlog(c, "abcdnosfna;rbg;abr") //this is invalid
	updateBlog(c, id)                 //this will update the blog to new information
}
