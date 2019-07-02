package main

import (
	"github.com/dannypsnl/redux/v2/rematch"
)

type AppState struct {
	Node    NodeInfo // info about current node
	Peers   []Peer   // list of peers for the node
	Current int      // currently selected peer
}

type AppModel struct {
	rematch.Reducer
	State AppState
}

func (m *AppModel) SetInfo(s AppState, node NodeInfo) AppState {
	s.Node = node
	return s
}

func (m *AppModel) Current(s AppState, peerIndex int) AppState {
	// NOTE Not sure if I should just ignore invalid values or panic
	if peerIndex >= 0 && peerIndex < len(s.Peers) {
		s.Current = peerIndex
	}
	return s
}

func (m *AppModel) Update(s AppState, peers []Peer) AppState {
	// The argument is a copy so we can modify it and return it
	s.Peers = peers
	// if not current peer is set use first one
	if s.Current == -1 && len(peers) > 0 {
		s.Current = 0
	}
	// if set but doesn't exist in the list move up
	if s.Current >= len(peers) {
		s.Current = len(peers) - 1
	}
	return s
}
