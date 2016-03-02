package orc

import (
	"errors"
	"fmt"
)

type NodeService interface {
	Register(string, string) error
}

type nodeService struct {
	nodes []node
}

func NewNodeService() *nodeService {
	return new(nodeService)
}

func (ns *nodeService) Register(name, addr string) error {
	for _, n := range ns.nodes {
		if n.address == addr {
			return errors.New("Node already exists")
		}
	}

	ns.nodes = append(ns.nodes, node{name: name, address: addr})
	fmt.Printf("Nodes: %#v\n", ns.nodes)
	return nil
}

type node struct {
	name, address string
}
