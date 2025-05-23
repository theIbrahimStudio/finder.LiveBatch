package batcher

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/theIbrahimStudio/finder.LiveBatch/backend"
	"github.com/theIbrahimStudio/finder.LiveBatch/config"
	"github.com/theIbrahimStudio/finder.LiveBatch/models"
)

type Batcher struct {
	config  config.Config
	queue   *RequestQueue
	backend backend.Backend
}

func NewBatcher(config config.Config, backend backend.Backend) *Batcher {
	return &Batcher{
		config:  config,
		queue:   NewRequestQueue(),
		backend: backend,
	}
}

func (b *Batcher) HandleRequest(w http.ResponseWriter, r *http.Request) {
	var req models.InferRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	respChan := make(chan models.InferResponse)
	b.queue.Enqueue(RequestWithResponse{Request: req, ResponseChan: respChan})
	resp := <-respChan
	json.NewEncoder(w).Encode(resp)
}

func (b *Batcher) Run() {
	ticker := time.NewTicker(time.Duration(b.config.MaxLatencyMs) * time.Millisecond)
	for range ticker.C {
		batch := b.queue.Dequeue(b.config.MaxBatchSize)
		if len(batch) == 0 {
			continue
		}
		requests := make([]models.InferRequest, len(batch))
		for i, req := range batch {
			requests[i] = req.Request
		}
		results := b.backend.Process(requests)
		for i, res := range results {
			batch[i].ResponseChan <- res
		}
	}
}
