package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
	"strings"
	"time"
)

type PeersState struct {
	c        *client
	list     []Peer
	selected *Peer
}

func NewPeersState(host string, port int) *PeersState {
	url := fmt.Sprintf("http://%s:%d", host, port)
	c, err := newClient(url)
	if err != nil {
		log.Panicln(err)
	}
	return &PeersState{c: c}
}

func (p *PeersState) FetchLoop(g *gocui.Gui) {
	for {
		select {
		case <-threadDone:
			return
		default:
			peers := p.Fetch()
			writePeers(g, peers)
			writePeerDetails(g, p.selected)
		}
		<-time.After(interval * time.Second)
	}
}

func (p *PeersState) Fetch() []Peer {
	peers, err := p.c.getPeers()
	if err != nil {
		log.Panicln(err)
	}
	p.list = peers
	if p.selected == nil {
		p.selected = &peers[0]
	}
	return peers
}

func writePeers(g *gocui.Gui, peers []Peer) {
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("main")
		if err != nil {
			return err
		}
		v.Clear()
		maxWidth, _ := g.Size()
		for _, peer := range peers {
			fmt.Fprintf(v, "%s\n", peer.AsTable(maxWidth))
		}
		return nil
	})
}

func writePeerDetails(g *gocui.Gui, peer *Peer) {
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("info")
		if err != nil {
			return err
		}
		v.Clear()
		fmt.Fprintf(v, strings.Repeat("%-8s: %s\n", 6),
			"Name", peer.Name,
			"ID", string(peer.Id),
			"Enode", peer.Enode,
			"Local", peer.Network.LocalAddress,
			"Remote", peer.Network.RemoteAddress,
			"Caps", strings.Join(peer.Caps, ", "))
		return nil
	})
}

func boolToString(v bool, yes string, no string) string {
	if v {
		return yes
	} else {
		return no
	}
}

func (p Peer) AsTable(maxWidth int) string {
	var id string
	if maxWidth > 160 {
		id = string(p.Id)
	} else {
		id = p.Id.String()
	}
	return fmt.Sprintf("%s ｜  %-15s ｜  %-21s ｜  %-7s ｜  %-8s",
		id, p.Name,
		p.Network.RemoteAddress,
		boolToString(p.Network.Trusted, "trusted", "normal"),
		boolToString(p.Network.Static, "static", "dynamic"))
}
