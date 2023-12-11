package main

import (
	"fmt"
	"net"
	"time"
	"sync"
)

const clientRequests = 99 // same as in client.go
const batchSize = 33
const splits = clientRequests/batchSize

func handleHttpRequest(conn net.Conn, wg *sync.WaitGroup) {
	defer conn.Close()
	defer wg.Done()

	response := "HTTP/1.1 200 OK"

	_, err := conn.Read(make([]byte, 1024))
	if err != nil {
		response = "HTTP/1.1 500 Internal Server Error"
	}

	conn.Write([]byte(response))
}

func dispatchRequestBatch(listener net.Listener, wg *sync.WaitGroup) {
	defer wg.Done()

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error when connecting with client: %s", err)
		return
	}

	wg.Add(1)
	go handleHttpRequest(conn, wg)
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	defer listener.Close()
	if err != nil {
		fmt.Println("Error when creating listener %s", err)
		return
	}
	
	wg := sync.WaitGroup{}
	startTime := time.Now().UnixNano()
	for i := 0; i < splits; i++ {
		for j := 0; j < batchSize; j++ {
			wg.Add(1)
			go dispatchRequestBatch(listener, &wg)
		}
	}
	wg.Wait()
	fmt.Println((time.Now().UnixNano() - startTime))
}