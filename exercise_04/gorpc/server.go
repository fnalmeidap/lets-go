package main

import (
	"gorpc/impl"
	"fmt"
	"net"
	"net/rpc"
)

const requests = 10000

func main() {
	api := new(impl.Api)

	server := rpc.NewServer()
	server.RegisterName("Api", api)

	listener, err := net.Listen("tcp", "localhost:8081")
	defer listener.Close()
	if err != nil {
		fmt.Println("Error creating listener:", err)
		return
	}

	server.Accept(listener)
}