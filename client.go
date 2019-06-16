package main

import (
	"github.com/ethereum/go-ethereum/rpc"
)

type client struct {
	rpcClient *rpc.Client
}

func newClient(url string) (*client, error) {
	rpcClient, err := rpc.Dial(url)
	if err != nil {
		return nil, err
	}
	return &client{rpcClient}, nil
}

func (c *client) getPeers() (interface{}, error) {
	var rval interface{}
	err := c.rpcClient.Call(&rval, "admin_peers")
	if err != nil {
		return nil, err
	}
	return rval, nil
}
