package p2p

import (
	"fmt"
	"log"
	"net"
	"sync"
)

// the remote node over tcp connection
type TCPPeer struct {
	conn net.Conn
	//if we dial and accept the connection, we are the outbound peer
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	mu            sync.RWMutex
	decoder       Decoder
	shakeHands    HandshakeFunc
	peers         map[net.Addr]Peer
}

func NewTCPTransport(listenAddress string) *TCPTransport {
	return &TCPTransport{
		shakeHands:    NOPHandShakeFunc,
		listenAddress: listenAddress,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	ln, err := net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}
	t.listener = ln
	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
		}
		go t.handleConn(conn)
	}
}

type Temp struct{}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)
	if err := t.shakeHands(peer); err != nil {
		fmt.Println("Error during handshake:", err)

	}

	//buf := new(bytes.Buffer)
	msg := &Temp{}
	for {
		if err := t.decoder.Decode(conn, msg); err != nil {
			log.Fatal("Error decoding message:", err)
			continue
		}
	}
}
