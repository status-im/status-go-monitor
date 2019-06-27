package main

import (
	"time"
)

func FetchLoop(client *StatusGoClient, state *State) {
	for {
		select {
		case <-threadDone:
			return
		default:
			state.Fetch(client)
		}
		<-time.After(interval * time.Second)
	}
}
