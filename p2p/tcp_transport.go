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

type TCPTransportOpts struct {
	ListenAddr    string
	HandshakeFunc HandshakeFunc
	Decoder       Decoder
}

type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener
	mu       sync.RWMutex
	decoder  Decoder
	peers    map[net.Addr]Peer
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	ln, err := net.Listen("tcp", t.ListenAddr)
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

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)
	if err := t.HandshakeFunc(peer); err != nil {
		fmt.Print("TCP handshake error", err)
		conn.Close()
		return
	}

	//buf := new(bytes.Buffer)
	msg := &Message{}
	for {
		if err := t.Decoder.Decode(conn, msg); err != nil {
			log.Fatal("Error decoding message:", err)
			continue
		}
		msg.From = conn.RemoteAddr()
		fmt.Print("Received message:", *msg)
	}
}
