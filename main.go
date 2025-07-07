package main

import (
	"fmt"
	"os"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
)

const privateKeyFile = "private.key"

func NewPeerID() crypto.PrivKey {
	priv, _, err := crypto.GenerateKeyPair(crypto.Ed25519, -1)
	if err != nil {
		panic(err)
	}

	id, err := peer.IDFromPublicKey(priv.GetPublic())
	if err != nil {
		panic(err)
	}

	fmt.Println("Peer ID:", id)

	return priv
}

func loadPrivateKey(file string) (crypto.PrivKey, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	priv, err := crypto.UnmarshalPrivateKey(data)
	if err != nil {
		return nil, err
	}

	return priv, nil
}

func savePrivateKey(file string, priv crypto.PrivKey) error {
	data, err := crypto.MarshalPrivateKey(priv)
	if err != nil {
		return err
	}

	return os.WriteFile(file, data, 0600)
}

func main() {
	
	priv, err := loadPrivateKey(privateKeyFile)

	if err != nil {
		fmt.Println("Generating new private key")

		priv, _, err = crypto.GenerateKeyPair(crypto.Ed25519, -1)
		if err != nil {
			panic(err)
		}

		err = savePrivateKey(privateKeyFile, priv)
		if err != nil {
			panic(err)
		}
	}

	id, err := peer.IDFromPublicKey(priv.GetPublic())
	if err != nil {
		panic(err)
	}

	fmt.Println("Peer ID:", id)
}
