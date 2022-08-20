package main

import (
	"context"
	"log"

	pb "github.com/Ansh5461/Go_gRPC/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("----createBlog was invoked----")

	blog := &pb.Blog{
		AuthorId: "Ansh",
		Title:    "My first blog",
		Content:  "Content of my first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error %v\n", err)
	}

	log.Printf("New blog has been created: %s\n", res.Id)
	return res.Id
}
