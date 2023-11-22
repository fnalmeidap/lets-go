package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const SERVER_PORT string = ":1997"

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	// todo: processar body
	fmt.Print("Received a POST request:\n")
	fmt.Print(string(body))

	w.Write([]byte("Message received!"))
}

func serverImpl(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		handlePostRequest(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}
}

func main() {
	http.HandleFunc("/", serverImpl)

	fmt.Println("Starting server on http://localhost" + SERVER_PORT)
	err := http.ListenAndServe(SERVER_PORT, nil)
	if err != nil {
		panic(err)
	}
}
