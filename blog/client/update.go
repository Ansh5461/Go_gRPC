package main

import (
	"context"
	"log"

	pb "github.com/Ansh5461/Go_gRPC/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("----updateBlog was invoked----")

	newblog := &pb.Blog{
		Id:       id,
		AuthorId: "Not Ansh",
		Title:    "A new title",
		Content:  "Content of the first blog, with some additions",
	}

	_, err := c.UpdateBlog(context.Background(), newblog)
	if err != nil {
		log.Fatalf("Error happened during updating: %v\n", err)
	}

	log.Println("Blog was updated")
}
