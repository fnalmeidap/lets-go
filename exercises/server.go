package main

import (
	"net/http"
)

const SERVER_PORT string = ":1997"

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// do something
}

func main() {
	http.HandleFunc("/", handleRequest)

	err := http.ListenAndServe(SERVER_PORT, nil)
	if err != nil {
		panic(err)
	}
}
