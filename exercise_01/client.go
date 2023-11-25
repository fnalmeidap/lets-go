package main

import (
	"bufio"
	"fmt"
	"net"
)

func sendHttpRequest(conn net.Conn) {
	request := "Hello!\n"
	conn.Write([]byte(request))
	// request := "GET / HTTP/1.1\r\n" +
	// 	"Host:  localhost:8080\r\n" +
	// 	"Connection: close\r\n" +
	// 	"\r\n"

	_, err := bufio.NewReader(conn).ReadString('\n')
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