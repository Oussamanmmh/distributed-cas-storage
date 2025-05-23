package main

import (
	"fmt"
	"log"

	"github.com/oussamanmmh/distributed-cas-storage/p2p"
)

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    "localhost:4000",
		HandshakeFunc: p2p.NOPHandShakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}
	tr := p2p.NewTCPTransport(tcpOpts)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatalf("Error starting TCP transport: %v", err)
	}
	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("Received message: %s\n", msg.Payload)
		}
	}()
	select {}
}
