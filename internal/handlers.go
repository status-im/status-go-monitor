package internal

import (
	G "github.com/jroimartin/gocui"
)

// Key handlers beyond this point
func (vm *ViewManager) Refresh(g *G.Gui, v *G.View) error {
	vm.Control.Fetch()
	return nil
}

func (vm *ViewManager) CursorUp(g *G.Gui, v *G.View) error {
	current := vm.Control.State.GetData().Current
	vm.Control.State.SetCurrentPeer(current - 1)
	return nil
}

func (vm *ViewManager) CursorDown(g *G.Gui, v *G.View) error {
	current := vm.Control.State.GetData().Current
	vm.Control.State.SetCurrentPeer(current + 1)
	return nil
}

func (vm *ViewManager) HandleDelete(g *G.Gui, v *G.View) error {
	handler := func() (err error) {
		currentPeer := vm.Control.State.GetCurrent()
		err = vm.Control.RemovePeer(currentPeer)
		return
	}
	if err := createConfirmView(vm, "Delete this peer?", handler); err != nil {
		return err
	}
	return nil
}

func (vm *ViewManager) HandleTrust(g *G.Gui, v *G.View) error {
	currentPeer := vm.Control.State.GetCurrent()
	if err := vm.Control.TrustPeer(currentPeer); err != nil {
		return err
	}
	return nil
}
