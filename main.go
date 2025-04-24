package main

import (
	"log"

	"github.com/oussamanmmh/distributed-cas-storage/p2p"
)

func main() {
	tr := p2p.NewTCPTransport("localhost:4000")
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatalf("Error starting TCP transport: %v", err)
	}
	select {}
}
