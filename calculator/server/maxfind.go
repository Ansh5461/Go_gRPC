package main

import (
	"io"
	"log"

	pb "github.com/Ansh5461/Go_gRPC/calculator/proto"
)

func (s *Server) MaxTillNow(stream pb.PrimeFind_MaxTillNowServer) error {

	//here we will implement an infinite loop for receving the information and sending it back
	log.Println("Function MaxTillNow was invoked")
	maxn := int64(0)

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Got an error while receiving request, reading client stream %v", err)
			break
		}

		x := res.Inp

		if x > maxn {
			maxn = x
			err = stream.Send(&pb.Result{
				Max: maxn,
			})

			if err != nil {
				log.Fatalf("Got error while sending it back %v", err)
			}
		}

	}

	return nil
}
