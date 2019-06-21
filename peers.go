package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
	"time"
)

func FetchPeers(c *client, g *gocui.Gui) {
	for {
		select {
		case <-threadDone:
			return
		default:
			peers, err := c.getPeers()
			if err != nil {
				log.Panicln(err)
			}
			WritePeers(g, peers)
		}
		<-time.After(interval * time.Second)
	}
}

func WritePeers(g *gocui.Gui, peers []Peer) {
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("main")
		if err != nil {
			return err
		}
		v.Clear()
		maxWidth, _ := g.Size()
		fmt.Fprintf(v, "%-15s | %-40s | %s\n", "Peer ID", "Name", "Protocols")
		for _, peer := range peers {
			fmt.Fprintf(v, "%s\n", peer.AsTable(maxWidth))
		}
		return nil
	})
}
