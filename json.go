package main

import "fmt"

type Peer struct {
	Enode     string            `json:"enode"`
	Id        string            `json:"id"`
	Name      string            `json:"na"`
	Caps      []string          `json:"caps"`
	Network   NetworkInfo       `json:"netrowkr"`
	Protocols map[string]string `json:"protocols"`
}

func (p Peer) String() string {
	return fmt.Sprintf("Peer(id=%s)", p.Id)
}

type NetworkInfo struct {
	LocalAddress  string `json:"localAddress"`
	RemoteAddress string `json:"remoteAddress"`
	Inbound       bool   `json:"inbound"`
	Trusted       bool   `json:"trusted"`
	Static        bool   `json:"static"`
}
