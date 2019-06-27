package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jroimartin/gocui"
)

type rcpResp map[string]interface{}

const host = "127.0.0.1"
const port = 8545
const interval = 5

var threadDone = make(chan struct{})

// TODO Add command line options
func main() {
	// Custom location for log messages
	clientLogFile, err := os.OpenFile("./app.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Panicln(err)
	}
	log.SetOutput(clientLogFile)

	// Core object for the Terminal UI
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	// Client necessary for doing RPC calls to status-go
	url := fmt.Sprintf("http://%s:%d", host, port)
	client, err := newClient(url)
	if err != nil {
		log.Panicln(err)
	}

	// Create a state wrapper.
	state := NewState(client)
	// Subscribe rendering method to state changes.
	state.Store.Subscribe(GenRenderFunc(g, state))

	mainView := &ViewController{
		Name:        "main",
		Title:       "Peers",
		Placeholder: "Loading peers...",
		Enabled:     true,
		Cursor:      true,
		Highlight:   true,
		Current:     true,
		SelFgColor:  gocui.ColorBlack,
		SelBgColor:  gocui.ColorGreen,
		State:       state,
		// corner positions
		TopLeft:  func(mx, my int) (int, int) { return 0, 0 },
		BotRight: func(mx, my int) (int, int) { return mx - 1, my / 2 },
	}
	// bindings defined separately so handlers can reference mainView
	mainView.Keybindings = []Binding{
		Binding{gocui.KeyCtrlC, gocui.ModNone, quit},
		Binding{gocui.KeyArrowUp, gocui.ModNone, mainView.CursorUp},
		Binding{gocui.KeyArrowDown, gocui.ModNone, mainView.CursorDown},
		Binding{'k', gocui.ModNone, mainView.CursorUp},
		Binding{'j', gocui.ModNone, mainView.CursorDown},
		Binding{gocui.KeyDelete, gocui.ModNone, mainView.HandleDelete},
		Binding{'d', gocui.ModNone, mainView.HandleDelete},
	}
	infoView := &ViewController{
		Name:        "info",
		Title:       "Details",
		Placeholder: "Loading details...",
		Enabled:     true,
		Wrap:        true,
		State:       state,
		// corner positions
		TopLeft:  func(mx, my int) (int, int) { return 0, (my / 2) + 1 },
		BotRight: func(mx, my int) (int, int) { return mx - 1, my - 1 },
	}
	// TODO Create a prompt view for user convirmations.

	views := []*ViewController{mainView, infoView}

	vm := ViewManager{g: g, views: views}

	g.SetManagerFunc(vm.Layout)

	// Start RPC calling routine for fetching peers periodically.
	go FetchLoop(state, interval)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	close(threadDone)
	return gocui.ErrQuit
}
