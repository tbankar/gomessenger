package main

import (
	"log"
	"net"

	handler "github.com/gomessenger/pkg/api"
	"github.com/gomessenger/pkg/proto"
	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

const (
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
	return
}
