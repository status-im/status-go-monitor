package main

import (
	"log"

	_ "github.com/dannypsnl/redux"
	"github.com/dannypsnl/redux/rematch"
	"github.com/dannypsnl/redux/store"
)

type State struct {
	Reducer     *Model
	Store       *store.Store
	updatePeers *rematch.Action
	setCurrent  *rematch.Action
}

func NewState() *State {
	// Generate the reducer from our model
	Reducer := &Model{
		State: PeersState{
			Peers:   make([]Peer, 0),
			Current: nil,
		},
	}
	// Instantiate the redux state from the reducer
	return &State{
		Reducer: Reducer,
		// Define the store
		Store: store.New(Reducer),
		// Define available reducers for the store
		updatePeers: Reducer.Action(Reducer.Update),
		setCurrent:  Reducer.Action(Reducer.Current),
	}
}

// Helpers for shorter calls
func (s *State) Update(peers []Peer) {
	s.Store.Dispatch(s.updatePeers.With(peers))
}
func (s *State) SetCurrent(peer *Peer) {
	s.Store.Dispatch(s.setCurrent.With(peer))
}
func (s *State) GetState() PeersState {
	return s.Store.StateOf(s.Reducer).(PeersState)
}
func (s *State) Fetch(client *StatusGoClient) {
	peers, err := client.getPeers()
	if err != nil {
		log.Panicln(err)
	}
	log.Printf("peers: %v\n", peers)
	ps := s.GetState()
	s.Update(peers)
	if ps.Current == nil {
		s.SetCurrent(&peers[0])
	}
}

func (s *State) Remove(client *StatusGoClient, peer Peer) error {
	success, err := client.removePeer(peer.Enode)
	if err != nil || success != true {
		log.Panicln(err)
	}
	s.Fetch(client)
	return nil
}
