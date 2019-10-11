package main

import (
	"fmt"
	"gomessenger/server/internal"
	"log"

	"github.com/streadway/amqp"
)

func stopOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
	}

}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@mqserver:5672/")
	if err != nil {
		log.Fatal("Error while connecting to RabbitMQ server", err)
	}

	defer conn.Close()
	ch, err := conn.Channel()
	stopOnError(err, "Failed to open channel")

	q, err := ch.QueueDeclare(
		"exec_server_rpc", // name
		false,             // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	stopOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	stopOnError(err, "Failed to register a consumer")

	waitForever := make(chan bool)
	go func() {
		for msg := range msgs {
			switch msg.Headers["action"] {
			case "create":
				var resp string
				err := internal.CreateUser(msg.Body)
				if err != nil {
					resp = err.Error()
				} else {
					resp = "1"
				}
				err = ch.Publish(
					"",
					msg.ReplyTo,
					false,
					false,
					amqp.Publishing{
						ContentType:   "text/plain",
						CorrelationId: msg.CorrelationId,
						Body:          []byte(resp),
					})
			case "login":
				err := internal.LoginUser(msg.Body)
				if err != nil {
					fmt.Println(err)
				}
			}
			msg.Ack(false)
		}
	}()
	fmt.Println("Waiting...")
	<-waitForever

}
