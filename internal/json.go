package internal

import "fmt"

type NodeInfo struct {
	Enode       string                  `json:"enode"`
	Name        string                  `json:"name"`
	ID          PeerID                  `json:"id"`
	ListenIp    string                  `json:"ip"`
	ListenAddr  string                  `json:"listenAddr"`
	ListenPorts NodeInfoPorts           `json:"ports"`
	Protocols   map[string]NodeProtocol `json:"protocols"`
}

type NodeInfoPorts struct {
	Discovery int `json:"discovery"`
	Listener  int `json:"listener"`
}

type NodeProtocol struct {
	MaxMessageSize int     `json:"maxMessageSize"`
	MinimumPoW     float32 `json:"minimumPoW"`
	Version        string  `json:"version"`
}

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

type PeerID string

// the ID is too long to display in full in most places
func (id PeerID) String() string {
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
