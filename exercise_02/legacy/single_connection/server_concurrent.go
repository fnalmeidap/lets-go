package main

import (
	"fmt"
	"net"
	"time"
	"sync"
)

const clientRequests = 100000 // same as in client.go
const batchSize = 5
const splits = clientRequests/batchSize

func handleHttpRequest(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	response := "HTTP/1.1 200 OK"

	_, err := conn.Read(make([]byte, 1024))
	if err != nil {
		response = "HTTP/1.1 500 Internal Server Error"
	}

	conn.Write([]byte(response))
}

func dispatchRequestBatch(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 0; j <= batchSize; j++ {
		wg.Add(1)
		go handleHttpRequest(conn, wg)
	}
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
	
	wg := sync.WaitGroup{}

	startTime := time.Now().UnixNano()
	for i := 0; i < splits; i++ {
		wg.Add(1)
		go dispatchRequestBatch(conn, &wg);
	}
	wg.Wait()
	fmt.Println((time.Now().UnixNano() - startTime))
}