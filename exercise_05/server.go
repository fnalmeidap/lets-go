package main

import (
	"encoding/json"
	"fmt"
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

	fmt.Println("here")

	ch, err := conn.Channel()
	checkError(err)
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		RequestQueue,
		false, // Durable
		false, // Delete when unused
		false, // Exclusive
		false, // No-wait
		nil,   // Arguments
	)
	checkError(err)

	msgs, err := ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		true, // nowait
		nil,
	)
	checkError(err)

	for msg := range msgs {
		response := "hello client"
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
