package main

import (
	"github.com/dannypsnl/redux/v2/rematch"
)

type PeersState struct {
	Peers   []Peer
	Current int
}

type PeersModel struct {
	rematch.Reducer
	State PeersState
}

func (m *PeersModel) Current(state PeersState, peerIndex int) PeersState {
	// NOTE Not sure if I should just ignore invalid values or panic
	if peerIndex >= 0 && peerIndex < len(state.Peers) {
		state.Current = peerIndex
	}
	return state
}

func (m *PeersModel) Update(state PeersState, peers []Peer) PeersState {
	// The argument is a copy so we can modify it and return it
	state.Peers = peers
	// if not current peer is set use first one
	if state.Current == -1 && len(peers) > 0 {
		state.Current = 0
	}
	// if set but doesn't exist in the list move up
	if state.Current >= len(peers) {
		state.Current = len(peers) - 1
	}
	return state
}
