
package main

import (
	"fmt"
	"net"
	"sync"
)

const requests = 10000

func handleHttpRequest(wg *sync.WaitGroup, conn *net.UDPConn, addr *net.UDPAddr, message string) {
	defer wg.Done()
	response := []byte(message)
	_, err := conn.WriteToUDP(response, addr)
	if err != nil {
		fmt.Println("Error sending response:", err)
	}
}

func main() {
	serverAddr, err := net.ResolveUDPAddr("udp", "localhost:8082")
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

	wg := sync.WaitGroup{}
	for i := 0; i < requests; i++ {
		wg.Add(1)
		response := "HTTP/1.1 200 OK"
		buffer := make([]byte, 1024)

		_, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			response = "HTTP/1.1 500 Internal Server Error"
		}

		go handleHttpRequest(&wg, conn, addr, response)
	}
	wg.Wait()
}
