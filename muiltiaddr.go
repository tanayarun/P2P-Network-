package main

import (
	"fmt"

	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
)

func create_multiaddr() {
	maddr, err := ma.NewMultiaddr("/ip4/127.0.0.1/tcp/8080")
	if err != nil {
		panic(err)
	}

	fmt.Println("Multiaddr: ", maddr)

	netAddr, err := manet.ToNetAddr(maddr)

	if err != nil {
		panic(err)
	}

	fmt.Println("Netaddr: ", netAddr.String())
}

func append_multiaddr() {
	maddr, err := ma.NewMultiaddr("/ip4/127.0.0.1/tcp/8080")
	if err != nil {
		panic(err)
	}

	wsaddr, err := ma.NewMultiaddr("/ws")
	if err != nil {
		panic(err)
	}

	finaladdr := maddr.Encapsulate(wsaddr)
	fmt.Println("Final Multiadd : ", finaladdr)
}
