package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTranspo(t *testing.T) {

	listenAddress := "localhost:4000"
	opts := TCPTransportOpts{
		ListenAddr:    listenAddress,
		HandshakeFunc: NOPHandShakeFunc,
		Decoder:       DefaultDecoder{},
	}
	tr := NewTCPTransport(opts)

	assert.Equal(t, listenAddress, tr.ListenAddr)
	assert.Nil(t, tr.ListenAndAccept())

}
