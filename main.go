package main

import (
	"fmt"
	"log"
	"time"
	//"github.com/kr/pretty"
	"github.com/jroimartin/gocui"
)

type rcpResp map[string]interface{}

const host = "127.0.0.1"
const port = 8545
const interval = 5

var threadDone = make(chan struct{})

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.Cursor = true

	g.SetManagerFunc(layout)

	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}

	url := fmt.Sprintf("http://%s:%d", host, port)
	c, err := newClient(url)
	if err != nil {
		log.Panicln(err)
	}

	// Start RPC calling routine
	go fetchPeers(c, g)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func fetchPeers(c *client, g *gocui.Gui) {
	for {
		select {
		case <-threadDone:
			return
		case <-time.After(interval * time.Second):
			peers, err := c.getPeers()
			if err != nil {
				log.Panicln(err)
			}
			writePeers(g, peers)
		}
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

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("main", 0, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.SetCursor(0, 0)
		fmt.Fprintln(v, "Loading peers...")
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	close(threadDone)
	return gocui.ErrQuit
}
