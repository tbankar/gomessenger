package api

import (
	"context"

	"github.com/gomessenger/pkg/proto"
)

type Server struct{}

func (s *Server) CreateUser(ctx context.Context, in *proto.CreateUserInput) *proto.CreateUserOutput {
	email := in.Email
	username := in.Username
	name := in.Name
}
