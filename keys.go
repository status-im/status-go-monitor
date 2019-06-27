package main

import (
	"github.com/jroimartin/gocui"
)

type Binding struct {
	Key     interface{} // so both gocui.Key and rune work
	Mod     gocui.Modifier
	Handler func(g *gocui.Gui, v *gocui.View) error
}

func (vc *ViewController) CursorUp(g *gocui.Gui, v *gocui.View) error {
	// TODO propper error handling?
	vc.State.SetCurrent(vc.State.GetState().Current - 1)
	return nil
}

func (vc *ViewController) CursorDown(g *gocui.Gui, v *gocui.View) error {
	// TODO propper error handling?
	vc.State.SetCurrent(vc.State.GetState().Current + 1)
	return nil
}

func (vc *ViewController) HandleDelete(g *gocui.Gui, v *gocui.View) error {
	currentPeer := vc.State.GetCurrent()
	err := vc.State.Remove(currentPeer)
	if err != nil {
		return err
	}
	return nil
}
