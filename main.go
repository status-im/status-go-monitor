package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	G "github.com/jroimartin/gocui"

	"github.com/status-im/status-go-monitor/internal"
)

type rcpResp map[string]interface{}

func main() {
	rpcAddr := flag.String("rpc-addr", "127.0.0.1", "IP address of the status-go RPC endpoint.")
	rpcPort := flag.Int("rpc-port", 8545, "TCP port of the status-go RPC endpoint. ")
	interval := flag.Int("interval", 3, "Interval in seconds for querying for list of peers.")
	flag.Parse()

	// Custom location for log messages
	clientLogFile, err := os.OpenFile("./app.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Panicln(err)
	}
	log.SetOutput(clientLogFile)

	// Core object for the Terminal UI
	g, err := G.NewGui(G.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	// Allow for handling of Escape key
	g.InputEsc = true
	defer g.Close()

	// Client necessary for doing RPC calls to status-go
	url := fmt.Sprintf("http://%s:%d", *rpcAddr, *rpcPort)
	client, err := internal.NewClient(url)
	if err != nil {
		log.Panicln(err)
	}

	// Create ViewManager without views to use in key bindings
	vm := internal.ViewManager{Gui: g}

	// Create a state wrapper.
	state := internal.NewState()

	// Create a state controller
	stateCtrl := &internal.StateController{
		State:  state,
		Client: client,
	}

	// Main view with list of peers
	mainView := &internal.ViewController{
		Name:        "main",
		Title:       "Peers",
		Placeholder: "Loading peers...",
		Cursor:      true,
		Highlight:   true,
		SelFgColor:  G.ColorBlack,
		SelBgColor:  G.ColorGreen,
		StateCtrl:   stateCtrl,
		// corner positions
		TopLeft:  func(mx, my int) (int, int) { return 0, 0 },
		BotRight: func(mx, my int) (int, int) { return mx - 1, my / 2 },
	}
	// bindings defined separately so handlers can reference mainView
	mainView.Keybindings = []internal.Binding{
		{G.KeyCtrlL, vm.Refresh},
		{'r', vm.Refresh},
		{G.KeyArrowUp, vm.CursorUp},
		{G.KeyArrowDown, vm.CursorDown},
		{'k', vm.CursorUp},
		{'j', vm.CursorDown},
		{G.KeyDelete, vm.HandleDelete},
		{'d', vm.HandleDelete},
		{'t', vm.HandleTrust},
	}
	// For viewing peer details
	infoView := &internal.ViewController{
		Name:        "info",
		Title:       "Details",
		Placeholder: "Loading details...",
		Wrap:        true,
		OnTop:       true,
		StateCtrl:   stateCtrl,
		// corner positions
		TopLeft:  func(mx, my int) (int, int) { return 0, (my / 2) + 1 },
		BotRight: func(mx, my int) (int, int) { return mx - 1, my - 1 },
	}

	// We attach views later so key bindings can access vm
	vm.Views = map[string]*internal.ViewController{
		"main": mainView,
		"info": infoView,
	}

	// Set the starting current view
	vm.Current = mainView.Name

	g.SetManagerFunc(vm.Layout)

	// Subscribe rendering method to state changes.
	state.Store.Subscribe(internal.GenRenderFunc(g, stateCtrl))

	// Start RPC calling routine for fetching peers periodically.
	go internal.FetchLoop(stateCtrl, *interval)

	if err := g.MainLoop(); err != nil && err != G.ErrQuit {
		log.Panicln(err)
	}
}
