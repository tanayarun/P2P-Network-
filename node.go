package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"

    "github.com/libp2p/go-libp2p"
)

func main() {
    host, err := libp2p.New()
    if err != nil {
        panic(err)
    }
    defer host.Close()

    fmt.Println("Listening on:")
    for _, addr := range host.Addrs() {
        fmt.Println(addr)
    }

    sigCh := make(chan os.Signal, 1)
    signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
    <-sigCh
}
