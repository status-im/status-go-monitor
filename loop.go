package main

import (
	"time"
)

func FetchLoop(state *State, interval int) {
	for {
		select {
		case <-threadDone:
			return
		default:
			state.Fetch()
		}
		<-time.After(time.Duration(interval) * time.Second)
	}
}
