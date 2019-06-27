package main

import (
	"github.com/dannypsnl/redux/rematch"
)

type PeersState struct {
	Peers   []Peer
	Current *Peer
}

type Model struct {
	rematch.Reducer
	State PeersState
}

type Todo struct {
	Title string
	Done  bool
}

func (todo *Model) Current(state PeersState, peer *Peer) PeersState {
	return PeersState{
		Peers:   state.Peers,
		Current: peer,
	}
}

func (todo *Model) Update(state PeersState, peers []Peer) PeersState {
	current := state.Current
	if state.Current == nil {
		current = &peers[0]
	}
	return PeersState{
		Peers:   peers,
		Current: current,
	}
}
