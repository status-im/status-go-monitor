package internal

import (
	"github.com/ethereum/go-ethereum/rpc"
)

type StatusGoClient struct {
	rpcClient *rpc.Client
}

func NewClient(url string) (*StatusGoClient, error) {
	rpcClient, err := rpc.Dial(url)
	if err != nil {
		return nil, err
	}
	return &StatusGoClient{rpcClient}, nil
}

func (c *StatusGoClient) nodeInfo() (info NodeInfo, err error) {
	err = c.rpcClient.Call(&info, "admin_nodeInfo")
	return
}

func (c *StatusGoClient) getPeers() (peers []Peer, err error) {
	err = c.rpcClient.Call(&peers, "admin_peers")
	return
}

func (c *StatusGoClient) removePeer(enode string) (success bool, err error) {
	err = c.rpcClient.Call(&success, "admin_removePeer", enode)
	return
}

func (c *StatusGoClient) trustPeer(enode string) (success bool, err error) {
	err = c.rpcClient.Call(&success, "shh_markTrustedPeer", enode)
	return
}
