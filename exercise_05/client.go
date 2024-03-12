package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	REQUESTS      = 10000
	ResponseQueue = "response_queue"
	RequestQueue  = "request_queue"
)

func checkError(err error) {
	if err != nil {
		log.Fatalf("%s", err)
	}
}

func sendHttpRequest(ch *amqp.Channel, message string) {
	request := message
	requestBytes, err := json.Marshal(request)
	checkError(err)

	err = ch.Publish(
		"",
		RequestQueue,
		false,
		false,
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: "10",
			ReplyTo:       ResponseQueue,
			Body:          requestBytes,
		},
	)
	checkError(err)
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	checkError(err)
	defer conn.Close()

	ch, err := conn.Channel()
	checkError(err)
	defer ch.Close()

	replyQueue, err := ch.QueueDeclare(
		ResponseQueue,
		false,
		false,
		false,
		false,
		nil,
	)
	checkError(err)

	msgs, err := ch.Consume(
		replyQueue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	checkError(err)

	for i := 0; i < REQUESTS; i++ {
		startTime := time.Now().UnixNano()

		request := "POST /path HTTP/1.1\n" +
			"Host: localhost:8081\n" +
			"Content-Type: text/plain\n" +
			"Content-Length: 18\n" +
			"Hello from client!"

		sendHttpRequest(ch, request)

		msg := <-msgs
		var response string
		err = json.Unmarshal(msg.Body, &response)
		checkError(err)

		fmt.Println(time.Now().UnixNano() - startTime)
	}
}
