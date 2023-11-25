package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func handleConnection(conn net.Conn) {
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error when reading client's request!")
		panic(err)
	}

	fmt.Printf(message)

	// response := "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\n" +
	// "\r\nMensagem recebida com sucesso!" +
	// "\r\n"
	response := "Message received!\n"
	conn.Write([]byte(response))
}

func main() {
	fmt.Println("TCP server running in :8080")

	listener, err := net.Listen("tcp", ":8080")
	defer listener.Close()
	if err != nil {
		fmt.Println("Error when creating listener.")
		panic(err)
	}

	conn, err := listener.Accept()
	defer conn.Close()
	if err != nil {
		fmt.Println("Error when connecting with client: %s", err)
		return
	}

	startTime := time.Now().UnixNano()
	for i := 0; i < 100; i++ {
		handleConnection(conn)
	}
	fmt.Println((time.Now().UnixNano() - startTime) / 1e6)
}