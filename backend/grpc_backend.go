package backend

import (
	"context"
	"log"
	"time"

	"github.com/theIbrahimStudio/finder.LiveBatch/models"
	pb "github.com/theIbrahimStudio/finder.LiveBatch/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCBackend struct {
	client pb.InferenceClient
}

func NewGRPCBackend(address string) *GRPCBackend {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC backend: %v", err)
	}
	client := pb.NewInferenceClient(conn)
	return &GRPCBackend{client: client}
}

func (g *GRPCBackend) Process(batch []models.InferRequest) []models.InferResponse {
	reqs := make([]*pb.InferRequest, len(batch))
	for i, r := range batch {
		reqs[i] = &pb.InferRequest{Input: r.Input}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := g.client.PredictBatch(ctx, &pb.BatchInferRequest{Requests: reqs})
	if err != nil {
		log.Printf("gRPC backend error: %v", err)
		responses := make([]models.InferResponse, len(batch))
		for i := range responses {
			responses[i] = models.InferResponse{Result: "error"}
		}
		return responses
	}

	outputs := make([]models.InferResponse, len(res.Responses))
	for i, r := range res.Responses {
		outputs[i] = models.InferResponse{Result: r.Result}
	}
	return outputs
}
