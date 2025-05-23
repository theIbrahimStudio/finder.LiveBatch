package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/theIbrahimStudio/finder.LiveBatch/backend"
	"github.com/theIbrahimStudio/finder.LiveBatch/batcher"
	"github.com/theIbrahimStudio/finder.LiveBatch/config"
)

func LoadConfig() config.Config {
	return config.Config{
		MaxBatchSize: 8,
		MaxLatencyMs: 50,
		ListenAddr:   ":8080",
	}
}

func main() {
	config := config.Load()
	fmt.Println(config)

	b := batcher.NewBatcher(config, backend.NewDummyBackend())
	go b.Run()

	http.HandleFunc("/infer", b.HandleRequest)
	log.Printf("LiveBatch listening on %s\n", config.ListenAddr)
	log.Fatal(http.ListenAndServe(config.ListenAddr, nil))
}
