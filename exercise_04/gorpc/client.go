package main

import (
	"gorpc/impl"
	"fmt"
	"net/rpc"
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

	fmt.Println(response)
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

	sendHttpRequest(client)
}