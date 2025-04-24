package p2p

//Peer is an interface represent remote node
type Peer interface {
	Close() error
}

//Handles the communication between peers
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}
