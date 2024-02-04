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

	buffer := make([]byte, 1024)
	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Error reading server's response.")
		panic(err)
	}

	fmt.Println("Response from server:", string(buffer[:n]))
}

func main() {
	serverAddr, err := net.ResolveUDPAddr("udp", "localhost:8081")
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

		fmt.Printf("RTT for request %d: %d ms\n", i + 1, (time.Now().UnixNano() - startTime))
	}
}