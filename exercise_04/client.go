package main

import (
	"fmt"
	"gorpc/impl"
	"net/rpc"
	"time"
)

const requests = 30000

func sendHttpRequest(client *rpc.Client) {
	request := impl.Request{Message: "POST /path HTTP/1.1\n" +
		"Host: localhost:8081\n" +
		"Content-Type: text/plain\n" +
		"Content-Length: 18\n" +
		"Hello from client!"}

	response := impl.Response{}

	err := client.Call("Api.Greet", request, &response)
	if err != nil {
		fmt.Println("Error calling server.")
		panic(err)
	}
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:8081")
	defer func(client *rpc.Client) {
		var err = client.Close()
		if err != nil {
			fmt.Println("Error closing stub connection:", err)
			return
		}
	}(client)
	if err != nil {
		fmt.Println("Error dialing server.")
		panic(err)
	}

	for i := 0; i < requests; i++ {
		startTime := time.Now().UnixNano()
		sendHttpRequest(client)
		fmt.Println(time.Now().UnixNano() - startTime)
	}
}
