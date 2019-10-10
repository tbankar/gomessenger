package api

import (
	"bytes"
	"encoding/json"
	"math/rand"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

func create_corrID() string {
	const bytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 16)
	for i := range b {
		b[i] = bytes[rand.Intn(len(bytes))]
	}
	return string(b)
}

func getConn() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("msngrserver:8443", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func CallCreateUser(userinfo *CreateInputReq, responseChann chan<- string) {
	conn, err := amqp.Dial("amqp://guest:guest@mqserver:5672/")
	if err != nil {
		responseChann <- err.Error()
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		responseChann <- err.Error()
	}
	defer ch.Close()
	// TODO: Decide server queue depending on user request hardcoding as of now as there is only one server
	q, err := ch.QueueDeclare(
		userinfo.Username+"_Reply",
		false,
		false,
		false,
		true,
		nil,
	)
	if err != nil {
		responseChann <- err.Error()
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		responseChann <- err.Error()
	}

	corrID := create_corrID()

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(userinfo)
	h := amqp.Table{ACTION: CREATE}
	err = ch.Publish(
		"",
		"exec_server_rpc",
		false,
		false,
		amqp.Publishing{
			ContentType:   "text/plain",
			Body:          b.Bytes(),
			CorrelationId: corrID,
			ReplyTo:       q.Name,
			Headers:       h,
		},
	)
	if err != nil {
		responseChann <- err.Error()
	}

	for m := range msgs {
		if corrID == m.CorrelationId {
			responseChann <- string(m.Body)
		}
	}
	return
}
