package main

import (
	"log"
	"net/http"

	"distributed-kv-store/internal/api"
	"distributed-kv-store/internal/consensus"
	"distributed-kv-store/internal/kvstore"
	"distributed-kv-store/internal/monitoring"
)

func main() {
	zab := consensus.NewZab()

	for i := 0; i < 3; i++ {
		node := kvstore.NewNode()
		zab.AddNode(node)
	}

	r := api.SetupRouter(zab)
	monitoring.SetupMonitoring()

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
