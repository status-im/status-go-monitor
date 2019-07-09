package internal

import (
	"time"

	"github.com/jroimartin/gocui"
)

// Thread ending channel
var threadDone = make(chan struct{})

func FetchLoop(s *StateController, interval int) {
	// Get the first peers fetch going sooner
	s.Fetch()
	// Then fetch every `interval` seconds
	for {
		select {
		case <-threadDone:
			return
		case <-time.After(time.Duration(interval) * time.Second):
			s.Fetch()
		}
	}
}
func QuitLoop(g *gocui.Gui, v *gocui.View) error {
	close(threadDone)
	return gocui.ErrQuit
}
