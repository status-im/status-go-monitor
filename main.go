package main

import (
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

	peers := NewPeersState(host, port)

	mainView := &ViewController{
		Name:        "main",
		Title:       "Peers",
		Placeholder: "Loading peers...",
		Cursor:      true,
		Highlight:   true,
		Current:     true,
		SelFgColor:  gocui.ColorBlack,
		SelBgColor:  gocui.ColorGreen,
		Data:        peers,
		TopLeft: func(mx, my int) (int, int) {
			return 0, 0
		},
		BotRight: func(mx, my int) (int, int) {
			return mx - 1, my / 2
		},
	}
	// bindings defined separately so handlers can reference mainView
	mainView.Keybindings = []Binding{
		Binding{gocui.KeyCtrlC, gocui.ModNone, quit},
		Binding{gocui.KeyArrowUp, gocui.ModNone, mainView.CursorUp},
		Binding{gocui.KeyArrowDown, gocui.ModNone, mainView.CursorDown},
	}
	infoView := &ViewController{
		Name:        "info",
		Title:       "Details",
		Placeholder: "Loading details...",
		TopLeft: func(mx, my int) (int, int) {
			return 0, my/2 + 1
		},
		BotRight: func(mx, my int) (int, int) {
			return mx - 1, my - 1
		},
	}

	views := []*ViewController{mainView, infoView}

	vm := NewViewManager(g, views)

	g.SetManagerFunc(vm.Layout)

	// Start RPC calling routine
	go peers.Fetch(g)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	close(threadDone)
	return gocui.ErrQuit
}
