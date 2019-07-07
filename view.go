package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

// Struct for more succint definitions of key bindings
type Binding struct {
	Key     interface{} // so both gocui.Key and rune work
	Mod     gocui.Modifier
	Handler func(g *gocui.Gui, v *gocui.View) error
}

// Default Gocui views arent granular enough
// so I'm adding a custom one to have more control.
type ViewController struct {
	Name        string
	Title       string
	Placeholder string
	Enabled     bool
	Wrap        bool
	Cursor      bool
	Current     bool
	Highlight   bool
	OnTop       bool
	TopLeft     func(int, int) (int, int)
	BotRight    func(int, int) (int, int)
	SelBgColor  gocui.Attribute
	SelFgColor  gocui.Attribute
	Keybindings []Binding
	// Extra field for view state. Might need different name.
	StateCtrl *StateController
}

// To combine all existing views into one
type ViewManager struct {
	g     *gocui.Gui
	views []*ViewController
}

func (m *ViewManager) Layout(g *gocui.Gui) error {
	mx, my := g.Size()

	for _, cfg := range m.views {
		if !cfg.Enabled {
			continue
		}
		x0, y0 := cfg.TopLeft(mx, my)
		x1, y1 := cfg.BotRight(mx, my)

		v, err := g.SetView(cfg.Name, x0, y0, x1, y1)

		// Some settings can be set only once
		if err == gocui.ErrUnknownView {
			cfg.SetKeybindings(g)
			if cfg.Cursor {
				v.SetCursor(0, 0)
			}
			if cfg.Placeholder != "" {
				fmt.Fprintln(v, cfg.Placeholder)
			}
		} else if err != nil {
			log.Panicln(err)
		}

		v.Title = cfg.Title
		v.Wrap = cfg.Wrap
		v.SelFgColor = cfg.SelFgColor
		v.SelBgColor = cfg.SelBgColor
		v.Highlight = cfg.Highlight

		if cfg.Current {
			g.SetCurrentView(cfg.Name)
		}
		if cfg.OnTop {
			g.SetViewOnTop(cfg.Name)
		}
	}
	return nil
}

func (v *ViewController) SetKeybindings(g *gocui.Gui) error {
	for _, b := range v.Keybindings {
		if err := g.SetKeybinding("", b.Key, b.Mod, b.Handler); err != nil {
			return err
		}
	}
	return nil
}

func (vc *ViewController) Refresh(g *gocui.Gui, v *gocui.View) error {
	// TODO proper error handling?
	vc.StateCtrl.Fetch()
	return nil
}

func (vc *ViewController) CursorUp(g *gocui.Gui, v *gocui.View) error {
	// TODO proper error handling?
	current := vc.StateCtrl.State.GetData().Current
	vc.StateCtrl.State.SetCurrentPeer(current - 1)
	return nil
}

func (vc *ViewController) CursorDown(g *gocui.Gui, v *gocui.View) error {
	// TODO proper error handling?
	current := vc.StateCtrl.State.GetData().Current
	vc.StateCtrl.State.SetCurrentPeer(current + 1)
	return nil
}

func (vc *ViewController) HandleDelete(g *gocui.Gui, v *gocui.View) error {
	currentPeer := vc.StateCtrl.State.GetCurrent()
	err := vc.StateCtrl.RemovePeer(currentPeer)
	if err != nil {
		return err
	}
	return nil
}

func (vc *ViewController) HandleTrust(g *gocui.Gui, v *gocui.View) error {
	currentPeer := vc.StateCtrl.State.GetCurrent()
	err := vc.StateCtrl.TrustPeer(currentPeer)
	if err != nil {
		return err
	}
	return nil
}
