package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

type Binding struct {
	Key     gocui.Key
	Mod     gocui.Modifier
	Handler func(g *gocui.Gui, v *gocui.View) error
}

var bindings = [...]Binding{
	Binding{gocui.KeyCtrlC, gocui.ModNone, quit},
	Binding{gocui.KeyArrowUp, gocui.ModNone, HandlerCursorUp},
	Binding{gocui.KeyArrowDown, gocui.ModNone, HandlerCursorDown},
}

func keybindings(g *gocui.Gui) error {
	for _, b := range bindings {
		if err := g.SetKeybinding("", b.Key, b.Mod, b.Handler); err != nil {
			return err
		}
	}
	return nil
}

func HandlerCursorUp(g *gocui.Gui, v *gocui.View) error {
	fmt.Println("UP")
	return nil
}

func HandlerCursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}
