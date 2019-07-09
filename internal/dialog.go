package internal

import (
	G "github.com/jroimartin/gocui"
)

const HELP string = `
    h / ? : Show help
      Esc : Quit current view
  C-c / q : Quit app
  C-l / r : Refresh view
   Up / k : Move up
 Down / j : Move down
  Del / d : Delete peer
        t : Trust peer
`

func createConfirmView(vm *ViewManager, text string, handler func() error) error {
	vc := &ViewController{
		Name:        "confirm",
		Title:       "Confirm (Escape to close)",
		Placeholder: text + " (y/n)",
		Highlight:   true,
		TopLeft:     func(mx, my int) (int, int) { return (mx / 2) - 15, (my / 2) - 1 },
		BotRight:    func(mx, my int) (int, int) { return (mx / 2) + 15, (my / 2) + 1 },
	}
	confirm := func(g *G.Gui, v *G.View) error {
		if err := handler(); err != nil {
			return err
		}
		vm.RemoveView(vc)
		return nil
	}
	destroy := func(g *G.Gui, v *G.View) error {
		vm.RemoveView(vc)
		return nil
	}
	vc.Keybindings = []Binding{
		{'y', confirm},
		{'n', destroy},
		{'q', destroy},
		{G.KeyEsc, destroy},
	}
	vm.AddView(vc)
	vm.Current = vc.Name
	return nil
}

func createHelpView(vm *ViewManager) error {
	vc := &ViewController{
		Name:        "help",
		Title:       "Help (Escape to close)",
		Placeholder: HELP,
		Highlight:   true,
		TopLeft:     func(mx, my int) (int, int) { return (mx / 2) - 17, (my / 2) - 5 },
		BotRight:    func(mx, my int) (int, int) { return (mx / 2) + 17, (my / 2) + 5 },
	}
	destroy := func(g *G.Gui, v *G.View) error {
		vm.RemoveView(vc)
		return nil
	}
	vc.Keybindings = []Binding{
		{'q', destroy},
		{G.KeyEsc, destroy},
	}
	vm.AddView(vc)
	vm.Current = vc.Name
	return nil
}
