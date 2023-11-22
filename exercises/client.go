package main

import (
	"fmt"
)

func createMessage() (json_data map[string]string) {
	json_data = make(map[string]string)
	json_data["key1"] = "value1"
	json_data["key2"] = "value2"

	return
}

func main() {
	response, err := http.Post(ENDPOINT, "application/json", bytes.NewBuffer(createMessage()))
	if err != nil {
	todo: handle?
	panic(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
	  // todo: handle?
	  panic(err)
	}

	doSomethingWithBody(body)
}
