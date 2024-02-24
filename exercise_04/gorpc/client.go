package main

import (
	"gorpc/api"
	"fmt"
	"net/rpc"
)

const requests = 10000

func sendHttpRequest(client *rpc.Client) {
	request := api.Request{message: "Hello from client!"}
	response := api.Response{}

	// (fnap): discarding response on purpose to maintain experiment factors
	err = client.Call("Server.Greet", request, &response)
	if err != nil {
		fmt.Println("Error reading server's response.")
		panic(err)
	}
}

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:8082")
	defer func(client *rpc.Client) {
		var err = client.Close()
		if err != nil {
			fmt.Println("Error closing stub connection:", err)
			return
		}
	}(client)
	if err != nil {
		fmt.Println("Error resolving server address:", err)
		return
	}


	sendHttpRequest(client)

}