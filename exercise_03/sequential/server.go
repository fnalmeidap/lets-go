
package main

import (
	"fmt"
	"net"
)

const requests = 10000

func handleHttpRequest(conn *net.UDPConn, addr *net.UDPAddr, message string) {
	response := []byte(message)
	_, err := conn.WriteToUDP(response, addr)
	if err != nil {
		fmt.Println("Error sending response:", err)
	}
}

func main() {
	fmt.Println("Starting UDP server!")

	serverAddr, err := net.ResolveUDPAddr("udp", "localhost:8081")
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	conn, err := net.ListenUDP("udp", serverAddr)
	defer conn.Close()
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}

	fmt.Println("Waiting for requests...")

	for i := 0; ; i++ {
		response := "HTTP/1.1 200 OK"
		buffer := make([]byte, 1024)

		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			response = "HTTP/1.1 500 Internal Server Error"
		}

		handleHttpRequest(conn, addr, response)
		
		fmt.Printf("Request %d from address %s:\n%s\n\n", i+1, addr, string(buffer[:n]))
	}
	fmt.Println("here\n")
}
