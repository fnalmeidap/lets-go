package main

import (
	"fmt"
	"net"
)

func sendHttpRequest(conn net.Conn) {
	request := "POST /path HTTP/1.1\n" +
				"Host: localhost:8000\n" +
				"Content-Type: text/plain\n" +
				"Content-Length: 18\n" +
				"Hello from client!"
	conn.Write([]byte(request))

	_, err := conn.Read(make([]byte, 1024))
	if err != nil {
		fmt.Println("Error when reading server's response!")
		panic(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	defer conn.Close()
	if err != nil {
		fmt.Println("Error when connecting with server: %s", err)
		return
	}

	for i := 0; i < 100; i++ {
		sendHttpRequest(conn)
	}
}