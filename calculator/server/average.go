package main

import (
	"io"
	"log"

	pb "github.com/Ansh5461/Go_gRPC/calculator/proto"
)

func (s *Server) Average(stream pb.PrimeFind_AverageServer) error {
	log.Println("Average function was invoked")
	//we will receive a stream of numbers and we have to return the average of that, so we need 2 variables
	//1 for total, and other for total instances
	total := float64(0)
	i := float64(0)

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.Outp{
				Res: total / i,
			})
		}

		if err != nil {
			log.Fatalf("Received error while receiving stream %v\n", err)
		}

		total += float64(req.Num)
		i += 1
	}

}
