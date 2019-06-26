package main

import (
	"github.com/jroimartin/gocui"
)

type Binding struct {
	Key     gocui.Key
	Mod     gocui.Modifier
	Handler func(g *gocui.Gui, v *gocui.View) error
}

var bindings = [...]Binding{
	Binding{gocui.KeyCtrlC, gocui.ModNone, quit},
	Binding{gocui.KeyArrowUp, gocui.ModNone, HandlerCursorDispenser(-1)},
	Binding{gocui.KeyArrowDown, gocui.ModNone, HandlerCursorDispenser(1)},
}

func keybindings(g *gocui.Gui) error {
	for _, b := range bindings {
		if err := g.SetKeybinding("", b.Key, b.Mod, b.Handler); err != nil {
			return err
		}
	}
	return nil
}

func HandlerCursorDispenser(mod int) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
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
}
