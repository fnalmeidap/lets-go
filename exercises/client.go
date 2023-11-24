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

	return
}

func main() {
	client := &http.Client{}
	m := Message{"Hello!"}
	req := createRequest(m)
	
	for i := 0; i < 100; i++ {
		response, err := client.Do(req)
		defer response.Body.Close()
		if err != nil {
			panic(err)
		}
		
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		
		fmt.Print(string(body))
	}
}