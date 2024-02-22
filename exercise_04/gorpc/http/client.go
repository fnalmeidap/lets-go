package main

import (
	"lets-go/exercise_04/gorpc/impl"
	"fmt"
	"net/rpc"
	"time"
)

const requests = 10000

func sendHttpRequest(client *rpc.Client) {
	request := impl.Request{message: "Hello from client!"}
	response := impl.Response{}

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

	for i := 0; i < requests; i++ {
		startTime := time.Now().UnixNano()
		sendHttpRequest(client)
		fmt.Println(time.Now().UnixNano() - startTime)
	}
}