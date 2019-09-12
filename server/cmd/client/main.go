package main

import (
	"context"
	"fmt"
	"log"

	"gomessenger/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8443", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := proto.NewMessengerServiceClient(conn)
	cuserOut, err := client.CreateUser(context.Background(), &proto.CreateUserInput{Username: "tushar", Name: " Tushar Bankar", Email: "tabankar@gmail.com"})
	if err != nil {
		log.Fatal(err)
	}
	if cuserOut.Retmessage != "" {
		fmt.Println(cuserOut.Retmessage)
	} else {
		fmt.Println("User created")
	}
}
