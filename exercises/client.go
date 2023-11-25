package main
// message := "POST / HTTP/1.1\r\n" +
// 	"Host: localhost:8080\r\n" +
// 	"Content-Type: application/json\r\n" +
// 	"Content-Length: " + fmt.Sprint(len(messageBody)) + "\r\n" +
// 	messageBody +
// 	"\r\n"

import (
	"bufio"
	"fmt"
	"net"
)

func sendHttpRequest(conn net.Conn) {
	defer conn.Close()

	message := "Hello!\n"
	conn.Write([]byte(message))

	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Printf("Erro ao ler a resposta do servidor: %s\n", err)
		return
	}

	fmt.Printf(response)
}

func main() {
	for {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			fmt.Printf("Erro ao conectar ao servidor: %s\n", err)
			return
		}

		sendHttpRequest(conn)
	}
}