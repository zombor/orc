package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/zombor/orc"
	"github.com/zombor/orc/pb"
)

type grpcBinding struct {
	orc.NodeService
}

func (b grpcBinding) RegisterNode(ctx context.Context, req *pb.RegisterNodeRequest) (*pb.NodeStatus, error) {
	if err := b.validateRegister(req); err != nil {
		return &pb.NodeStatus{Ok: false}, err
	}

	err := b.NodeService.Register(req.Name, req.Address)

	if err == nil {
		// Dial the node address
		// Send `w` command

		conn, err := grpc.Dial(req.Address, grpc.WithInsecure())

		if err != nil {
			return &pb.NodeStatus{Ok: false}, err
		}

		client := pb.NewNodeClient(conn)
		res, err := client.RunCommand(
			context.Background(),
			&pb.RunCommandRequest{
				Command: "w",
				Params:  map[string]string{},
			},
		)

		if err != nil {
			return &pb.NodeStatus{Ok: false}, err
		} else {
			fmt.Println(res.Message)
		}
	}

	return &pb.NodeStatus{Ok: err == nil}, err
}

func (b grpcBinding) validateRegister(req *pb.RegisterNodeRequest) error {
	if req.Name == "" || req.Address == "" {
		var err registerNodeReqErr
		if req.Name == "" {
			err.Name = "Name is required"
		}
		if req.Address == "" {
			err.Address = "Address is required"
		}

		return err
	}

	return nil
}

type registerNodeReqErr struct {
	Name, Address string
}

func (err registerNodeReqErr) Error() string {
	var errs []string
	if err.Name != "" {
		errs = append(errs, err.Name)
	}
	if err.Address != "" {
		errs = append(errs, err.Address)
	}
	return strings.Join(errs, ", ")
}
