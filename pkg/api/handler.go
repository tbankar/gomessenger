package api

import (
	"context"

	"github.com/gomessenger/internal/datastore"
	"github.com/gomessenger/pkg/proto"
)

type Server struct{}

func (s *Server) CreateUser(ctx context.Context, in *proto.CreateUserInput) (*proto.CreateUserOutput, error) {
	//c := datastore.UserDetails{"Tushar_Bankar1",Useremail: in.Email, Name: in.Name, UserID: "Tushar_Bankar1"}
	c := datastore.UserDetails{"Tushar_Bankar1", in.Email, in.Username, in.Name}
	output, err := datastore.SendCreateUser(c)
	return &proto.CreateUserOutput{Retmessage: output}, nil

}
