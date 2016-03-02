package main

import (
	"golang.org/x/net/context"

	"github.com/zombor/orc"
	"github.com/zombor/orc/pb"
)

type grpcBinding struct {
	orc.CommandService
}

func (b grpcBinding) RunCommand(ctx context.Context, req *pb.RunCommandRequest) (*pb.RunCommandResponse, error) {
	command := req.Command
	params := req.Params

	msg, err := b.CommandService.Run(command, params)

	return &pb.RunCommandResponse{Message: msg, Ok: err == nil}, err
}
