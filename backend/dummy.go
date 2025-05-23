package backend

import (
	"time"

	"github.com/theIbrahimStudio/finder.LiveBatch/models"
)

type Backend interface {
	Process([]models.InferRequest) []models.InferResponse
}

type DummyBackend struct{}

func NewDummyBackend() *DummyBackend {
	return &DummyBackend{}
}

func (d *DummyBackend) Process(batch []models.InferRequest) []models.InferResponse {
	time.Sleep(10 * time.Millisecond)
	responses := make([]models.InferResponse, len(batch))
	for i := range batch {
		responses[i] = models.InferResponse{Result: "ok"}
	}
	return responses
}
