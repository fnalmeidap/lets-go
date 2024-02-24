package main

import (
	"gorpc/impl" // see README.md
	"fmt"
	"net/rpc"
	"time"
)

const requests = 10000

func sendHttpRequest(client *rpc.Client) {
	request := impl.Request{Message: "Hello from client!"}
	response := impl.Response{ }

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