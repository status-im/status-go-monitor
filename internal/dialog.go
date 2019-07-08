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
		{'u', destroy},
		{G.KeyEsc, destroy},
	}
	vm.AddView(vc)
	vm.Current = vc.Name
	return nil
}
