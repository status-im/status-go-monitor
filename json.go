package main

import "fmt"

type Peer struct {
	Enode     string            `json:"enode"`
	Id        peerId            `json:"id"`
	Name      string            `json:"name"`
	Caps      []string          `json:"caps"`
	Network   NetworkInfo       `json:"network"`
	Protocols map[string]string `json:"protocols"`
}

func (p Peer) String() string {
	return fmt.Sprintf("Peer(id=%s)", p.Id)
}

type peerId string

// the ID is too long to display in full in most places
func (id peerId) String() string {
	return fmt.Sprintf("%s...%s",
		string(id[:6]),
		string(id[len(id)-6:]))
}

type NetworkInfo struct {
	LocalAddress  string `json:"localAddress"`
	RemoteAddress string `json:"remoteAddress"`
	Inbound       bool   `json:"inbound"`
	Trusted       bool   `json:"trusted"`
	Static        bool   `json:"static"`
}
