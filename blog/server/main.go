package main

import (
	"context"
	"log"
	"net"

	pb "github.com/Ansh5461/Go_gRPC/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

//first we will create a client that will connecct to the mongoDB database, from this client
//we will be able to get the databases, and from this database we will access collection that we call
//blog

var collection *mongo.Collection

var addr string = "0.0.0.0:50051"

//Server struct is having comment just because vs says so
type Server struct {
	pb.BlogServiceServer
}

func main() {

	//docker is binding mongodb
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	//with this client we can connect to mongoDB

	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//now we will be getting the database and from the database we will be getting the collection
	//this collection will look for a database named blog DB,
	collection = (*mongo.Collection)(client.Database("blogdb").Collection("blog"))

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on due to: %v\n ", err)
	}

	log.Printf("Listening on %v\n", addr)

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v\n", err)
	}

	//now we will declare a mongo connection, which will help us to do everything with mongoDB
}
