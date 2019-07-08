package internal

import (
	"errors"
	"fmt"
	"log"

	G "github.com/jroimartin/gocui"
)

// Binding for more succint definitions of key bindings
type Binding struct {
	Key     interface{} // so both gocui.Key and rune work
	Handler func(g *G.Gui, v *G.View) error
}

// ViewController for more control, gocui.View is too bare
type ViewController struct {
	Name        string
	Title       string
	Placeholder string
	Wrap        bool
	Cursor      bool
	Highlight   bool
	OnTop       bool
	TopLeft     func(int, int) (int, int)
	BotRight    func(int, int) (int, int)
	SelBgColor  G.Attribute
	SelFgColor  G.Attribute
	Keybindings []Binding
	StateCtrl   *StateController
}

// To combine all existing views into one
type ViewManager struct {
	Gui     *G.Gui
	Views   map[string]*ViewController
	Current string
}

func (vm *ViewManager) Layout(g *G.Gui) error {
	// Allow for handling of Escape key
	g.InputEsc = true

	// Add global key-bindings for help
	help := func(g *G.Gui, v *G.View) error {
		return createHelpView(vm)
	}
	// Add global key-bindings for quitting
	globalBindings := []Binding{
		{G.KeyCtrlC, QuitLoop},
		{'q', QuitLoop},
		{'h', help},
		{'?', help},
	}
	if err := SetKeybindings("", globalBindings, g); err != nil {
		return err
	}

	// Configure active views
	for _, vc := range vm.Views {
		if _, err := vm.configureView(g, vc); err != nil {
			log.Panicln(err)
		}
	}
	// Remove views we don't know
	for _, v := range g.Views() {
		if _, ok := vm.Views[v.Name()]; !ok {
			g.DeleteKeybindings(v.Name())
			g.DeleteView(v.Name())
		}
	}
	return nil
}

func (vm *ViewManager) configureView(g *G.Gui, vc *ViewController) (*G.View, error) {
	mx, my := g.Size()
	// Calculate dimensions of new view
	x0, y0 := vc.TopLeft(mx, my)
	x1, y1 := vc.BotRight(mx, my)

	// Create the view
	v, err := g.SetView(vc.Name, x0, y0, x1, y1)

	// Some settings can be set only once
	if err == G.ErrUnknownView {
		if err := SetKeybindings(vc.Name, vc.Keybindings, g); err != nil {
			return nil, err
		}
		if vc.Cursor {
			g.Cursor = true
			v.SetCursor(0, 0)
		}
		if vc.Placeholder != "" {
			fmt.Fprintln(v, vc.Placeholder)
		}
		if vc.OnTop {
			g.SetViewOnTop(vc.Name)
		}
	} else if err != nil {
		return nil, err
	}

	v.Title = vc.Title
	v.Wrap = vc.Wrap
	v.SelFgColor = vc.SelFgColor
	v.SelBgColor = vc.SelBgColor
	v.Highlight = vc.Highlight

	if vc.Name == vm.Current {
		g.SetCurrentView(vc.Name)
	}
	return v, nil
}

func SetKeybindings(name string, bindings []Binding, g *G.Gui) error {
	for _, b := range bindings {
		err := g.SetKeybinding(name, b.Key, G.ModNone, b.Handler)
		if err != nil {
			return err
		}
	}
	return nil
}

func (vm *ViewManager) GetViewCtrl(v *G.View) (*ViewController, error) {
	for _, vc := range vm.Views {
		if vc.Name == v.Name() {
			return vc, nil
		}
	}
	return nil, errors.New("failed to find ViewController")
}

func (vm *ViewManager) AddView(vc *ViewController) {
	vm.Views[vc.Name] = vc
}

func (vm *ViewManager) RemoveView(vc *ViewController) {
	delete(vm.Views, vc.Name)
}
