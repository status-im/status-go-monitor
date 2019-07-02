package main

import (
	"time"
)

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
