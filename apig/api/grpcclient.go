package api

import (
	"context"

	"gomessenger/server/pkg/proto"

	"google.golang.org/grpc"
)

func getConn(host string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(host+":8443", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func CallCreateUser(userinfo *InputReq, host string, uCreated chan bool, errChann chan error) {
	conn, err := getConn(host)

	if err != nil {
		errChann <- err
	}
	defer conn.Close()
	client := proto.NewMessengerServiceClient(conn)
	created, err := client.CreateUser(context.Background(), &proto.CreateUserInput{Username: userinfo.Username, Name: userinfo.UserFullname, Email: userinfo.UserEmail, Password: userinfo.Password})

	if err != nil {
		errChann <- err
	} else if created.Res {
		uCreated <- true
	}
}
