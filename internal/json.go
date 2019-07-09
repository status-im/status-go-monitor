package internal

import "fmt"

// NodeInfo is a response from admin_nodeInfo RPC call
type NodeInfo struct {
	Enode       string                  `json:"enode"`
	Name        string                  `json:"name"`
	ID          PeerID                  `json:"id"`
	ListenIP    string                  `json:"ip"`
	ListenAddr  string                  `json:"listenAddr"`
	ListenPorts NodeInfoPorts           `json:"ports"`
	Protocols   map[string]NodeProtocol `json:"protocols"`
}

// NodeInfoPorts is part of the NodeInfo struct
type NodeInfoPorts struct {
	Discovery int `json:"discovery"`
	Listener  int `json:"listener"`
}

// NodeProtocol is part of the NodeInfo struct
type NodeProtocol struct {
	MaxMessageSize int     `json:"maxMessageSize"`
	MinimumPoW     float32 `json:"minimumPoW"`
	Version        string  `json:"version"`
}

// Peer is a struct from a list response to admin_peers RPC call
type Peer struct {
	Enode     string            `json:"enode"`
	ID        PeerID            `json:"id"`
	Name      string            `json:"name"`
	Caps      []string          `json:"caps"`
	Network   NetworkInfo       `json:"network"`
	Protocols map[string]string `json:"protocols"`
}

func (p Peer) String() string {
	return fmt.Sprintf("Peer(ID=%s)", p.ID)
}

// PeerID is just a string, but has its own type for formatting
type PeerID string

// the ID is too long to display in full in most places
func (id PeerID) String() string {
	return fmt.Sprintf("%s...%s",
		string(id[:6]),
		string(id[len(id)-6:]))
}

// NetworkInfo is part of the Peer struct
type NetworkInfo struct {
	LocalAddress  string `json:"localAddress"`
	RemoteAddress string `json:"remoteAddress"`
	Inbound       bool   `json:"inbound"`
	Trusted       bool   `json:"trusted"`
	Static        bool   `json:"static"`
}
