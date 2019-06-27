package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
	"strings"
	"time"
)

type PeersState struct {
	c    *client
	list []Peer
}

func NewPeersState(host string, port int) *PeersState {
	url := fmt.Sprintf("http://%s:%d", host, port)
	c, err := newClient(url)
	if err != nil {
		log.Panicln(err)
	}
	return &PeersState{c: c}
}

func (p *PeersState) Fetch(g *gocui.Gui) {
	for {
		select {
		case <-threadDone:
			return
		default:
			peers, err := p.c.getPeers()
			if err != nil {
				log.Panicln(err)
			}
			p.list = peers
			writePeers(g, peers)
		}
		<-time.After(interval * time.Second)
	}
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
	return fmt.Sprintf("%s｜ %-15s｜ %-21s｜ %-7s｜ %-8s｜ %s",
		id, p.Name,
		p.Network.RemoteAddress,
		boolToString(p.Network.Trusted, "trusted", "normal"),
		boolToString(p.Network.Static, "static", "dynamic"),
		strings.Join(p.Caps, ", "))
}