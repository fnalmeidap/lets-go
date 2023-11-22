package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

func createMessage() {
  // todo: do something
}

func main() {
  response, err := http.Get(ENDPOINT)
  if err != nil {
    // todo: handle?
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
