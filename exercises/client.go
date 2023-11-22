package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Message struct {
	content string
}

func createMessage() (m Message) {
	m = Message{"Message content!"}

	return
}

func createRequest(m Message) (req *http.Request) {
	endpoint := "http://localhost:1997"
	buf, err := json.Marshal(m.content)
	if err != nil {
		panic(err)
	}

	req, err = http.NewRequest("POST", endpoint, bytes.NewBuffer(buf))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	return
}

func useBody(body []byte) {
	// todo: do something with body
	fmt.Print(string(body))
}

func main() {
	client := &http.Client{}
	m := createMessage()
	req := createRequest(m)

	response, err := client.Do(req)
	if err != nil {
		//todo: handle?
		panic(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		// todo: handle?
		panic(err)
	}

	useBody(body)
}
