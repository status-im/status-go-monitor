package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
	"time"
)

type rcpResp map[string]interface{}

const host = "127.0.0.1"
const port = 8545
const interval = 3

var threadDone = make(chan struct{})

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

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

var idx int

func writePeers(g *gocui.Gui, peers interface{}) {
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("main")
		if err != nil {
			return err
		}
		v.Clear()
		idx++
		fmt.Fprintf(v, "idx: %v\n", idx)
		return nil
	})
}

func keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	return nil
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("main", 0, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Loading peers...")
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	close(threadDone)
	return gocui.ErrQuit
}
