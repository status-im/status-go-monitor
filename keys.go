package main

import (
	"github.com/jroimartin/gocui"
)

type Binding struct {
	Key     gocui.Key
	Mod     gocui.Modifier
	Handler func(g *gocui.Gui, v *gocui.View) error
}

func (vc *ViewController) CursorUp(g *gocui.Gui, v *gocui.View) error {
	return MoveCursor(-1, vc, g, v)
}

func (vc *ViewController) CursorDown(g *gocui.Gui, v *gocui.View) error {
	return MoveCursor(1, vc, g, v)
}

func MoveCursor(mod int, vc *ViewController, g *gocui.Gui, v *gocui.View) error {
	if v == nil {
		return nil
	}
	cx, cy := v.Cursor()
	// get peers
	ps := vc.State.(*PeersState)
	peers := ps.list
	// Don't go beyond available list of peers
	if cy+mod >= len(peers) || cy+mod < 0 {
		return nil
	}
	// update currently selected peer in the list
	ps.selected = &peers[cy+mod]
	writePeerDetails(g, ps.selected)
	if err := v.SetCursor(cx, cy+mod); err != nil {
		if mod == -1 {
			return nil
		}
		ox, oy := v.Origin()
		if err := v.SetOrigin(ox, oy+mod); err != nil {
			return err
		}
	}
	return nil
}
