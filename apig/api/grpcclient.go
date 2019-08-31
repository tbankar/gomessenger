package api

import (
	"context"

	"gomessenger/pkg/proto"

	"google.golang.org/grpc"
)

func getConn(host string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(host+":8443", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func CallCreateUser(userinfo *InputReq, host string) (string, error) {
	conn, err := getConn(host)
	if err != nil {
		return "", err
	}
	defer conn.Close()
	client := proto.NewMessengerServiceClient(conn)
	cuserOut, err := client.CreateUser(context.Background(), &proto.CreateUserInput{Username: userinfo.Username, Name: userinfo.UserFullname, Email: userinfo.UserEmail})
	if err != nil {
		return "", err
	} else if cuserOut.Retmessage != "" {
		return cuserOut.Retmessage, nil
	} else {
		return "", nil
	}
}
