package main

import (
	"github.com/dannypsnl/redux/rematch"
)

type PeersState struct {
	Peers   []Peer
	Current int
}

type Model struct {
	rematch.Reducer
	State PeersState
}

func (m *Model) Current(state PeersState, peerIndex int) PeersState {
	return PeersState{
		Peers:   state.Peers,
		Current: peerIndex,
	}
}

func (m *Model) Update(state PeersState, peers []Peer) PeersState {
	current := state.Current
	// if not current peer is set use first one
	if state.Current == -1 && len(peers) > 0 {
		current = 0
	}
	// if set but doesn't exist in the list move up
	if state.Current >= len(peers) {
		current = len(peers) - 1
	}
	return PeersState{
		Peers:   peers,
		Current: current,
	}
}