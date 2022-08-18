package main

import (
	pb "github.com/Ansh5461/Go_gRPC/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//BlogItem is to save it to mongodb, and convert it to Blog
type BlogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}

//it will be an object where we will be defining how it will be stored in mongo DB, using bson
//bson is just a variant of json

//we need to transform blogitem inside the message of type blog

func documentToBlog(data *BlogItem) *pb.Blog {
	//hex is converting id to string
	return &pb.Blog{
		Id:       data.ID.Hex(),
		AuthorId: data.AuthorID,
		Title:    data.Title,
		Content:  data.Content,
	}
}
