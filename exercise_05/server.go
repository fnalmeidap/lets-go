package main

import (
	"encoding/json"
	"log"

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

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	checkError(err)
	defer conn.Close()

	ch, err := conn.Channel()
	checkError(err)
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		RequestQueue,
		false,
		false,
		false,
		false,
		nil,
	)
	checkError(err)

	msgs, err := ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		true,
		nil,
	)
	checkError(err)

	for {
		msg := <-msgs

		var request string
		err = json.Unmarshal(msg.Body, &request)
		checkError(err)

		response := "HTTP/1.1 200 OK"
		responseBytes, err := json.Marshal(response)
		checkError(err)

		err = ch.Publish(
			"",
			msg.ReplyTo,
			false,
			false,
			amqp.Publishing{
				ContentType:   "text/plain",
				CorrelationId: msg.CorrelationId,
				Body:          responseBytes,
			},
		)
		checkError(err)
	}
}
