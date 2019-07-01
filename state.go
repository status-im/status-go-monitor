package main

import (
	"log"

	"github.com/dannypsnl/redux/v2/rematch"
	"github.com/dannypsnl/redux/v2/store"
)

// This might need renaming, since it also contains the Client.
// I need the client to make the RPC calls.
type State struct {
	Reducer     *PeersModel
	Store       *store.Store
	Client      *StatusGoClient
	updatePeers *rematch.Action
	setCurrent  *rematch.Action
}

func NewState(client *StatusGoClient) *State {
	// Generate the reducer from our model.
	Reducer := &PeersModel{
		State: PeersState{
			Peers:   make([]Peer, 0),
			Current: -1, // Should mean non selected.
		},
	}
	// Instantiate the redux state from the reducer.
	return &State{
		Reducer: Reducer,
		// Define the store.
		Store: store.New(Reducer),
		// Client for RPC calls.
		Client: client,
		// Define available reducers for the store.
		updatePeers: Reducer.Action(Reducer.Update),
		setCurrent:  Reducer.Action(Reducer.Current),
	}
}

// Helpers for shorter calls.
func (s *State) Update(peers []Peer) {
	s.Store.Dispatch(s.updatePeers.With(peers))
}
func (s *State) GetCurrent() *Peer {
	state := s.GetState()
	if state.Current == -1 {
		return nil
	}
	return &state.Peers[state.Current]
}
func (s *State) SetCurrent(index int) {
	s.Store.Dispatch(s.setCurrent.With(index))
}
func (s *State) GetState() PeersState {
	return s.Store.StateOf(s.Reducer).(PeersState)
}

// For fetching current state of peers from status-go
func (s *State) Fetch() {
	peers, err := s.Client.getPeers()
	if err != nil {
		log.Panicln(err)
	}
	ps := s.GetState()
	s.Update(peers)
	if ps.Current == -1 {
		s.SetCurrent(0)
	}
}

// For removing a selected peer from connected to status-go
func (s *State) Remove(peer *Peer) error {
	success, err := s.Client.removePeer(peer.Enode)
	if err != nil || success != true {
		log.Panicln(err)
	}
	s.Fetch()
	return nil
}
