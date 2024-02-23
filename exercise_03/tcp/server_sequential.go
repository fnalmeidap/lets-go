package main

import (
	"fmt"
	"net"
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
	listener, err := net.Listen("tcp", ":8081")
	defer listener.Close()
	if err != nil {
		fmt.Println("Error creating listener:", err)
		return
	}

	conn, err := listener.Accept()
	defer conn.Close()
	if err != nil {
		fmt.Println("Error connecting with client:", err)
		return
	}

	for i := 0; i < requests; i++ {
		handleHttpRequest(conn)
	}
}