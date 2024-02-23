package main

import (
	"fmt"
	"net"
	"sync"
)

const requests = 10000

func handleHttpRequest(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

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
		fmt.Println("Error when creating listener:", err)
		return
	}
	
	conn, err := listener.Accept()
	defer conn.Close()
	if err != nil {
		fmt.Println("Error when connecting with client:", err)
		return
	}

	wg := sync.WaitGroup{}
	for i := 0; i < requests; i++ {
		wg.Add(1)
		go handleHttpRequest(conn, &wg)
	}
	wg.Wait()
}
