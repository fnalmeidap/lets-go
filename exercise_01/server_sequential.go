package main

import (
	"fmt"
	"net"
	"time"
)

func handleHttpRequest(conn net.Conn) {
	response := "HTTP/1.1 200 OK"

	_, err := conn.Read(make([]byte, 1024))
	if err != nil {
		response = "HTTP/1.1 500 Internal Server Error"
	}

	conn.Write([]byte(response))
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	defer listener.Close()
	if err != nil {
		fmt.Println("Error when creating listener %s", err)
		return
	}

	conn, err := listener.Accept()
	defer conn.Close()
	if err != nil {
		fmt.Println("Error when connecting with client: %s", err)
		return
	}

	startTime := time.Now().UnixNano()
	for i := 0; i < 100; i++ {
		handleHttpRequest(conn)
	}
	fmt.Println((time.Now().UnixNano() - startTime))
}