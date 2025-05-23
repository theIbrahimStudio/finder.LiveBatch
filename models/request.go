package models

type InferRequest struct {
	Input string `json:"input"`
}

type InferResponse struct {
	Result string `json:"result"`
}
