package main

import (
	"time"
)

func FetchLoop(state *State) {
	for {
		select {
		case <-threadDone:
			return
		default:
			state.Fetch()
		}
		<-time.After(interval * time.Second)
	}
}
