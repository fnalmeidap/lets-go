package main

import (
	"gorpc/api"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

const requests = 10000

func handleHttpRequest(conn net.Conn) {
	response := "HTTP/1.1 200 OK"

	_, err := conn.Read(make([]byte, 1024))
	if err != nil {
		response = "HTTP/1.1 500 Internal Server Error"
	}

	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Error sending response:", err)
	}
}

func main() {
	api := new(api.Api)

	server := rpc.NewServer()
	server.RegisterName("Api", api)

	server.HandleHTTP("/", "/debug")

	listener, err := net.Listen("tcp", ":8081")
	defer listener.Close()
	if err != nil {
		fmt.Println("Error creating listener:", err)
		return
	}

	http.Serve(listener, nil)
}