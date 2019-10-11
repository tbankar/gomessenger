package api

import (
	"bytes"
	"encoding/json"
	"gomessenger/common"
	"math/rand"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

func createCorrID() string {
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

	corrID := createCorrID()

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

func LogToDB(userLogin *LoginInputReq) {
	conn, err := amqp.Dial("amqp://guest:guest@mqserver:5672/")
	if err != nil {
		common.PopulateError(err, "Error while dialing a connection to MQServer")
	}
	defer conn.Close()
	ch, err := conn.Channel()
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"exec_server_rpc", // name
		false,             // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		common.PopulateError(err, "Error while Decalring a queue")
	}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(userLogin)
	h := amqp.Table{ACTION: LOGIN}
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        b.Bytes(),
			Headers:     h,
		},
	)
	return

}
