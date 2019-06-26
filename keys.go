package main

import (
	"log"

	"github.com/jroimartin/gocui"
)

type Binding struct {
	Key     gocui.Key
	Mod     gocui.Modifier
	Handler func(g *gocui.Gui, v *gocui.View) error
}

func (vc *ViewController) CursorUp(g *gocui.Gui, v *gocui.View) error {
	return MoveCursor(-1, v)
}

func (vc *ViewController) CursorDown(g *gocui.Gui, v *gocui.View) error {
	peers := vc.Data.(*PeersState).list
	_, cy := v.Cursor()
	// Don't go beyond available list of peers
	if cy+1 >= len(peers) {
		return nil
	}
	return MoveCursor(1, v)
}

func MoveCursor(mod int, v *gocui.View) error {
	if v == nil {
		return nil
	}
	_, my := v.Size()
	cx, cy := v.Cursor()
	log.Printf("my: %d, cx: %d, cy: %d", my, cx, cy)
	if cy+mod < 0 || cy+mod == my {
		return nil
	}
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
