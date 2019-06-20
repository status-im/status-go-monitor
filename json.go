package main

import (
	"fmt"
	"strings"
)

type Peer struct {
	Enode     string            `json:"enode"`
	Id        peerId            `json:"id"`
	Name      string            `json:"name"`
	Caps      []string          `json:"caps"`
	Network   NetworkInfo       `json:"netrowkr"`
	Protocols map[string]string `json:"protocols"`
}

func (p Peer) String() string {
	return fmt.Sprintf("Peer(id=%s)", p.Id)
}

type peerId string

func (p Peer) AsTable(maxWidth int) string {
	var id string
	if maxWidth > 50 {
		id = string(p.Id)
	} else {
		id = p.Id.String()
	}
	return fmt.Sprintf("%15s | %30s | %s", id, p.Name, strings.Join(p.Caps, ", "))
}

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
