package main

import (
	"github.com/dannypsnl/redux/v2/rematch"
)

type AppData struct {
	Node    NodeInfo // info about current node
	Peers   []Peer   // list of peers for the node
	Current int      // currently selected peer
}

type AppModel struct {
	rematch.Reducer
	State AppData
}

func (m *AppModel) SetInfo(s AppData, node NodeInfo) AppData {
	s.Node = node
	return s
}

func (m *AppModel) Current(s AppData, peerIndex int) AppData {
	// NOTE Not sure if I should just ignore invalid values or panic
	if peerIndex >= 0 && peerIndex < len(s.Peers) {
		s.Current = peerIndex
	}
	return s
}

func (m *AppModel) Update(s AppData, peers []Peer) AppData {
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
