package main

import (
	"fmt"
	"net"
	"time"
	"sync"
)

const requests = 100
const batches = 5
const btSize = requests/batches

func handleBatchHttpRequest(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < btSize; i++ {
		response := "HTTP/1.1 200 OK"
		_, err := conn.Read(make([]byte, 1024))
		if err != nil {
			response = "HTTP/1.1 500 Internal Server Error"
		}
	
		conn.Write([]byte(response))
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8081")
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
	for i := 0; i < batches; i++ {
		wg.Add(1)
		go handleBatchHttpRequest(conn, &wg)
	}
	wg.Wait()
	fmt.Println((time.Now().UnixNano() - startTime))
}