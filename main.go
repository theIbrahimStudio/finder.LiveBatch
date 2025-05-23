package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/theIbrahimStudio/finder.LiveBatch/backend"
	"github.com/theIbrahimStudio/finder.LiveBatch/batcher"
	"github.com/theIbrahimStudio/finder.LiveBatch/config"
)

func main() {
	config := config.Load()
	fmt.Println(config)

	var b *batcher.Batcher

	if config.UseGRPC {
		b = batcher.NewBatcher(config, backend.NewGRPCBackend("localhost:50051"))
	} else {
		b = batcher.NewBatcher(config, &backend.DummyBackend{})
	}

	b.Run()
	http.HandleFunc("/infer", b.HandleRequest)
	log.Printf("LiveBatch listening on %s\n", config.ListenAddr)
	log.Fatal(http.ListenAndServe(config.ListenAddr, nil))
}
