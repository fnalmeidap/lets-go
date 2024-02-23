package main

import (
	"fmt"
	"net"
	"time"
)

const requests = 10000

func sendHttpRequest(conn net.Conn) {
	request := "POST /path HTTP/1.1\n" +
				"Host: localhost:8000\n" +
				"Content-Type: text/plain\n" +
				"Content-Length: 18\n" +
				"Hello from client!"
				
	_, err := conn.Write([]byte(request))
	if err != nil {
		fmt.Println("Error sending message.")
		panic(err)
	}

	_, err = conn.Read(make([]byte, 1024))
	if err != nil {
		fmt.Println("Error reading server's response!")
		panic(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8081")
	defer conn.Close()
	if err != nil {
		fmt.Println("Error when connecting with server: %s", err)
		return
	}

	for i := 0; i < requests; i++ {
		startTime := time.Now().UnixNano()
		sendHttpRequest(conn)
		fmt.Println(time.Now().UnixNano() - startTime)
	}
}