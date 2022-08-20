package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Ansh5461/Go_gRPC/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//CreateBlog was literally defined in proto file, yet its showing warning, dont know why
func (s *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	//we will create a blogitem instance and we will try to insert it in database
	log.Printf("Create blog was invoked with: %v\n", in)

	data := BlogItem{
		AuthorID: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	res, err := collection.InsertOne(ctx, data)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v\n", err),
		)
	}
	//if no error then means we have inserted the item in database

	//it's of type interface, we will cast it into objectid
	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			"Error while converting to OID",
		)
	}

	//hex is converting object id into string
	return &pb.BlogId{
		Id: oid.Hex(),
	}, nil
}
