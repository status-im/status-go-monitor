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

func (c *client) getPeers() ([]Peer, error) {
	peers := make([]Peer, 0)
	err := c.rpcClient.Call(&peers, "admin_peers")
	if err != nil {
		return nil, err
	}
	return peers, nil
}

func (c *client) removePeer(enode string) (bool, error) {
	var rval bool
	err := c.rpcClient.Call(&rval, "admin_removePeer", enode)
	if err != nil {
		return false, err
	}
	return rval, nil
}
