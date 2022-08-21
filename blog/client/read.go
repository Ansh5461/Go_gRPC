package main

import (
	"context"
	"log"

	pb "github.com/Ansh5461/Go_gRPC/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("----function readBlog was invoked----")

	req := &pb.BlogId{Id: id}

	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Printf("Error happened while printing : %v\n", err)
	}

	log.Printf("Blog was read : %v\n", res)
	return res
}
