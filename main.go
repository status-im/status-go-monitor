package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
	"os"
)

type rcpResp map[string]interface{}

const host = "127.0.0.1"
const port = 8545
const interval = 5

var threadDone = make(chan struct{})

func main() {
	clientLogFile, err := os.OpenFile("./app.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Panicln(err)
	}
	log.SetOutput(clientLogFile)

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SelFgColor = gocui.ColorGreen
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
	go FetchPeers(c, g)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("main", 0, 0, maxX-1, maxY/2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.SelFgColor = gocui.ColorBlack
		v.SelBgColor = gocui.ColorGreen
		v.Title = "Peers"
		v.Highlight = true
		v.SetCursor(0, 0)
		g.SetCurrentView("main")
		fmt.Fprintln(v, "Loading peers...")
	}
	if v, err := g.SetView("info", 0, maxY/2+1, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Details"
		fmt.Fprintln(v, "Loading details...")
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	close(threadDone)
	return gocui.ErrQuit
}
