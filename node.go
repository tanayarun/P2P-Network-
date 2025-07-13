package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/libp2p/go-libp2p"
)

func NodeStart() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	node, err := libp2p.New(
		
		libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"),
	)
	if err != nil {
		panic(err)
	}
	defer node.Close()

	fmt.Println("Peer ID:", node.ID())
	fmt.Println("Listen addresses:", node.Addrs())

	
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	fmt.Println("Shutting down node...")
}
