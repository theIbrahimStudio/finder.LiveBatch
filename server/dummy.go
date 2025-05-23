package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/theIbrahimStudio/finder.LiveBatch/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedInferenceServer
}

func (s *server) PredictBatch(ctx context.Context, req *pb.BatchInferRequest) (*pb.BatchInferResponse, error) {
	responses := make([]*pb.InferResponse, len(req.Requests))
	for i, r := range req.Requests {
		responses[i] = &pb.InferResponse{Result: "echo:" + r.Input}
	}
	return &pb.BatchInferResponse{Responses: responses}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterInferenceServer(s, &server{})

	fmt.Println("Dummy gRPC Inference Server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
