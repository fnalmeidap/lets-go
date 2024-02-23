package main

import (
	"fmt"
	"net"
	"time"
)

const requests = 10000

func sendHttpRequest(conn *net.UDPConn) {
	request := "POST /path HTTP/1.1\n" +
				"Host: localhost:8081\n" +
				"Content-Type: text/plain\n" +
				"Content-Length: 18\n" +
				"Hello from client!"

	_, err := conn.Write([]byte(request))
	if err != nil {
		fmt.Println("Error sending message.")
		panic(err)
	}

	_, _, err = conn.ReadFromUDP(make([]byte, 1024))
	if err != nil {
		fmt.Println("Error reading server's response.")
		panic(err)
	}
}

func main() {
	serverAddr, err := net.ResolveUDPAddr("udp", "localhost:8082")
	if err != nil {
		fmt.Println("Error resolving server address:", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	defer conn.Close()
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}

	for i := 0; i < requests; i++ {
		startTime := time.Now().UnixNano()
		sendHttpRequest(conn)
		fmt.Println(time.Now().UnixNano() - startTime)
	}
}