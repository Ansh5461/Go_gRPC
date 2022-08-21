package main

import (
	"context"
	"log"

	pb "github.com/Ansh5461/Go_gRPC/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//ReadBlog function
func (s *Server) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {
	log.Printf("Function ReadBlog was invoked with : %v\n", in)

	//now we will get the id inside in(blogid) type and transform it to object id type
	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot pass id",
		)
	}

	data := &BlogItem{}
	//bson.M is a map data type in bson, and we are giving it oid, means, we are filtering by id, and value of id is oid
	filter := bson.M{"_id": oid}

	res := collection.FindOne(ctx, filter)

	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			"No blog with given ID",
		)
	}

	return documentToBlog(data), nil
}
