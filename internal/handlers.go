package internal

import (
	G "github.com/jroimartin/gocui"
)

// Key handlers beyond this point
func (vm *ViewManager) Refresh(g *G.Gui, v *G.View) error {
	vc, err := vm.GetViewCtrl(v)
	if err != nil {
		return err
	}
	vc.StateCtrl.Fetch()
	return nil
}

func (vm *ViewManager) CursorUp(g *G.Gui, v *G.View) error {
	vc, err := vm.GetViewCtrl(v)
	if err != nil {
		return err
	}
	current := vc.StateCtrl.State.GetData().Current
	vc.StateCtrl.State.SetCurrentPeer(current - 1)
	return nil
}

func (vm *ViewManager) CursorDown(g *G.Gui, v *G.View) error {
	vc, err := vm.GetViewCtrl(v)
	if err != nil {
		return err
	}
	current := vc.StateCtrl.State.GetData().Current
	vc.StateCtrl.State.SetCurrentPeer(current + 1)
	return nil
}

func (vm *ViewManager) HandleDelete(g *G.Gui, v *G.View) error {
	vc, err := vm.GetViewCtrl(v)
	if err != nil {
		return err
	}
	currentPeer := vc.StateCtrl.State.GetCurrent()
	if err := vc.StateCtrl.RemovePeer(currentPeer); err != nil {
		return err
	}
	return nil
}

func (vm *ViewManager) HandleTrust(g *G.Gui, v *G.View) error {
	vc, err := vm.GetViewCtrl(v)
	if err != nil {
		return err
	}
	currentPeer := vc.StateCtrl.State.GetCurrent()
	if err := vc.StateCtrl.TrustPeer(currentPeer); err != nil {
		return err
	}
	return nil
}

func (vm *ViewManager) HandleEscape(g *G.Gui, v *G.View) error {
	vc, err := vm.GetViewCtrl(v)
	if err != nil {
		return err
	}
	vc.OnTop = false
	return nil
}

func (vm *ViewManager) ShowHelp(g *G.Gui, v *G.View) error {
	vc, err := vm.GetViewCtrl(v)
	if err != nil {
		return err
	}
	vc.OnTop = true
	return nil
}
