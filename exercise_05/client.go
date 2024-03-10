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

func sendHttpRequest(ch *amqp.Channel) {
	request := "oi"
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

	//TODO: consume from reply queue
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
		false, // Durable
		false, // Delete when unused
		false, // Exclusive
		true,  // No-wait
		nil,   // Arguments
	)
	checkError(err)

	ch.Consume(
		replyQueue.Name,
		"",
		true,
		false,
		false,
		true, // nowait
		nil,
	)
	checkError(err)

	for i := 0; i < REQUESTS; i++ {
		startTime := time.Now().UnixNano()
		sendHttpRequest(ch)
		fmt.Println(time.Now().UnixNano() - startTime)
	}
}
