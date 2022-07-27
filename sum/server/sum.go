package main

import (
	"context"
	"log"

	pb "github.com/Ansh5461/Go_gRPC/sum/proto"
)

func (s *Server) Calculate(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Calculate from server was invoked by %v", in)
	return &pb.SumResponse{
		Res: in.Num1 + in.Num2,
	}, nil
}
