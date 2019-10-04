package api

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func CallCreateUser(userinfo *CreateInputReq, published chan bool, errChann chan<- error) {
	conn, err := amqp.Dial("amqp://guest:guest@172.17.0.3:5672/")
	if err != nil {
		errChann <- err
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		errChann <- err
	}
	defer ch.Close()
	// TODO: Decide server queue depending on user request hardcoding as of now as there is only one server
	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		false,
		true,
		nil,
	)
	if err != nil {
		errChann <- err
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
		errChann <- err
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
		errChann <- err
	}

	for m := range msgs {
		if corrID == m.CorrelationId {
			fmt.Println(string(m.Body))
		}
	}

	/*	client := proto.NewMessengerServiceClient(conn)
		created, err := client.CreateUser(context.Background(), &proto.CreateUserInput{Username: userinfo.Username, Name: userinfo.UserFullname, Email: userinfo.UserEmail, Password: userinfo.Password})

		if err != nil {
			errChann <- err
		} else if created.Res {
			uCreated <- true
		}*/
	return
}
