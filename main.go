package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jroimartin/gocui"

	"github.com/status-im/status-go-monitor/internal"
)

type rcpResp map[string]interface{}

const host = "127.0.0.1"
const port = 8545
const interval = 3

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
	client, err := internal.NewClient(url)
	if err != nil {
		log.Panicln(err)
	}

	// Create a state wrapper.
	state := internal.NewState()

	// Create a state controller
	stateCtrl := &internal.StateController{
		State:  state,
		Client: client,
	}

	// Subscribe rendering method to state changes.
	state.Store.Subscribe(internal.GenRenderFunc(g, stateCtrl))

	mainView := &internal.ViewController{
		Name:        "main",
		Title:       "Peers",
		Placeholder: "Loading peers...",
		Enabled:     true,
		Cursor:      true,
		Highlight:   true,
		Current:     true,
		SelFgColor:  gocui.ColorBlack,
		SelBgColor:  gocui.ColorGreen,
		StateCtrl:   stateCtrl,
		// corner positions
		TopLeft:  func(mx, my int) (int, int) { return 0, 0 },
		BotRight: func(mx, my int) (int, int) { return mx - 1, my / 2 },
	}
	// bindings defined separately so handlers can reference mainView
	mainView.Keybindings = []internal.Binding{
		{gocui.KeyCtrlC, gocui.ModNone, internal.QuitLoop},
		{gocui.KeyArrowUp, gocui.ModNone, mainView.CursorUp},
		{gocui.KeyArrowDown, gocui.ModNone, mainView.CursorDown},
		{'r', gocui.ModNone, mainView.Refresh},
		{gocui.KeyCtrlL, gocui.ModNone, mainView.Refresh},
		{'k', gocui.ModNone, mainView.CursorUp},
		{'j', gocui.ModNone, mainView.CursorDown},
		{gocui.KeyDelete, gocui.ModNone, mainView.HandleDelete},
		{'d', gocui.ModNone, mainView.HandleDelete},
		{'t', gocui.ModNone, mainView.HandleTrust},
	}
	infoView := &internal.ViewController{
		Name:        "info",
		Title:       "Details",
		Placeholder: "Loading details...",
		Enabled:     true,
		Wrap:        true,
		StateCtrl:   stateCtrl,
		// corner positions
		TopLeft:  func(mx, my int) (int, int) { return 0, (my / 2) + 1 },
		BotRight: func(mx, my int) (int, int) { return mx - 1, my - 1 },
	}
	// TODO Create a prompt view for user convirmations.

	views := []*internal.ViewController{mainView, infoView}

	vm := internal.ViewManager{Gui: g, Views: views}

	g.SetManagerFunc(vm.Layout)

	// Start RPC calling routine for fetching peers periodically.
	go internal.FetchLoop(stateCtrl, interval)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
