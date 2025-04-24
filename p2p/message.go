package p2p

import "net"

//Any arbitrary data that con be sent over the network
type Message struct {
	From    net.Addr
	Payload []byte
}
