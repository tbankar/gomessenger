package api

import (
	"bytes"
	"encoding/json"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

func getConn() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("msngrserver:8443", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func CallCreateUser(userinfo *CreateInputReq, uCreated chan bool, errChann chan error) {
	conn, err := amqp.Dial("amqp://guest:guest@172.17.0.2:5672/")
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
		"server1",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		errChann <- err
	}
	b := new(bytes.Buffer)
	m := make(map[string]CreateInputReq)
	m[CREATE] = *userinfo
	json.NewEncoder(b).Encode(m)
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/json",
			Body:        b.Bytes(),
		},
	)
	if err != nil {
		errChann <- err
	}

	/*	client := proto.NewMessengerServiceClient(conn)
		created, err := client.CreateUser(context.Background(), &proto.CreateUserInput{Username: userinfo.Username, Name: userinfo.UserFullname, Email: userinfo.UserEmail, Password: userinfo.Password})

		if err != nil {
			errChann <- err
		} else if created.Res {
			uCreated <- true
		}*/
}
