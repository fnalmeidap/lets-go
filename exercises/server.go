package main

import (
	"fmt"
	"net/http"
)

const SERVER_PORT string = ":1997"

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// do something
}

func main() {
	http.HandleFunc("/", handleRequest)

	if err := http.ListenAndServe(SERVER_PORT, nil); err != nil {
		panic(err)
	}
}
