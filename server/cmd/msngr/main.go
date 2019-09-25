package main

import (
	"fmt"
	"gomessenger/server/internal"
	"log"

	"github.com/streadway/amqp"
)

const (
	SERVER = "server1"
)

func stopOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
	}

}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@172.17.0.3:5672/")
	if err != nil {
		log.Fatalf("Error while connecting to RabbitMQ server", err)
	}

	defer conn.Close()
	ch, err := conn.Channel()
	stopOnError(err, "Failed to open channel")

	q, err := ch.QueueDeclare(
		SERVER, // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
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
				internal.CreateUser(msg.Body)
			}
		}
	}()
	fmt.Println("Waiting...")
	<-waitForever

}

/*const (
	GRPCPORT = ":8443"
)

func main() {
	lis, err := net.Listen("tcp", GRPCPORT)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer lis.Close()
	grpcServer := grpc.NewServer()
	proto.RegisterMessengerServiceServer(grpcServer, &handler.Server{})
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
	return*/
