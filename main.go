package main

import (
	"fmt"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
)

func main() {
	priv, _, err := crypto.GenerateKeyPair(crypto.Ed25519, -1)
	if err != nil {
		panic(err)
	}

	id, err := peer.IDFromPublicKey(priv.GetPublic())
	if err != nil {
		panic(err)
	}

	fmt.Println("Peer ID:", id)
}
