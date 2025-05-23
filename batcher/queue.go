package batcher

import "github.com/theIbrahimStudio/finder.LiveBatch/models"

type RequestWithResponse struct {
	Request      models.InferRequest
	ResponseChan chan models.InferResponse
}

type RequestQueue struct {
	queue []RequestWithResponse
}

func NewRequestQueue() *RequestQueue {
	return &RequestQueue{queue: []RequestWithResponse{}}
}

func (q *RequestQueue) Enqueue(req RequestWithResponse) {
	q.queue = append(q.queue, req)
}

func (q *RequestQueue) Dequeue(max int) []RequestWithResponse {
	if len(q.queue) == 0 {
		return nil
	}
	n := max
	if len(q.queue) < max {
		n = len(q.queue)
	}
	batch := q.queue[:n]
	q.queue = q.queue[n:]
	return batch
}
