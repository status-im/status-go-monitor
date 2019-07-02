package main

import (
	"time"
)

func FetchLoop(state *AppState, interval int) {
	// Get the first peers fetch going sooner
	state.Fetch()
	// Then fetch every `interval` seconds
	for {
		select {
		case <-threadDone:
			return
		case <-time.After(time.Duration(interval) * time.Second):
			state.Fetch()
		}
	}
}
