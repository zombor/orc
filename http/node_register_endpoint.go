package http

import (
	"golang.org/x/net/context"

	"github.com/go-kit/kit/endpoint"

	"github.com/zombor/ork"
)

func MakeRegisterNodeEndpoint(svc ork.NodeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(registerNodeRequest)
		err := svc.Register(req.Name, req.Address)
		if err != nil {
			return registerNodeResponse{""}, err
		}
		return registerNodeResponse{"OK"}, nil
	}
}
