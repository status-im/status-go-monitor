package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

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
	// extra field for view state
	State interface{}
}

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
		if err == nil && err != gocui.ErrUnknownView {
			return err
		}
		v.Title = cfg.Title
		v.Wrap = cfg.Wrap
		v.SelFgColor = cfg.SelFgColor
		v.SelBgColor = cfg.SelBgColor
		v.Highlight = cfg.Highlight

		if cfg.Cursor {
			v.SetCursor(0, 0)
		}
		if cfg.Current {
			g.SetCurrentView(cfg.Name)
		}
		if cfg.OnTop {
			g.SetViewOnTop(cfg.Name)
		}
		fmt.Fprintln(v, cfg.Placeholder)
		cfg.SetKeybindings(g)
	}
	return nil
}

func (v *ViewController) SetKeybindings(g *gocui.Gui) error {
	for _, b := range v.Keybindings {
		// IDEA: I can pass a method instead of a function here
		if err := g.SetKeybinding("", b.Key, b.Mod, b.Handler); err != nil {
			return err
		}
	}
	return nil
}
