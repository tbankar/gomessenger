package api

import (
	"context"

	"gomessenger/server/internal/datastore"
	"gomessenger/proto"
)

type Server struct{}

func (s *Server) CreateUser(ctx context.Context, in *proto.CreateUserInput) (*proto.CreateUserOutput, error) {
	c := datastore.UserDetails{UserID: in.Username, Useremail: in.Email, Username: in.Username, Name: in.Name, Password: in.Password}
	var c1 datastore.DstoreOps = &c
	ok, err := c1.CreateUser()
	if err != nil {
		return nil, err
	}
	return &proto.CreateUserOutput{Res: ok}, nil

}
