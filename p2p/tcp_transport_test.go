package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTranspo(t *testing.T) {

	listenAddress := "localhost:4000"
	tr := NewTCPTransport(listenAddress)

	assert.Equal(t, listenAddress, tr.listenAddress)
	assert.Nil(t, tr.ListenAndAccept())
	select {}
}
