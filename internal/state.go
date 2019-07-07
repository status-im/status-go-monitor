package internal

import (
	"github.com/dannypsnl/redux/v2/rematch"
	"github.com/dannypsnl/redux/v2/store"
)

// This might need renaming, since it also contains the Client.
// I need the client to make the RPC calls.
type AppState struct {
	Reducer        *AppModel
	Store          *store.Store
	setNodeInfo    *rematch.Action
	updatePeers    *rematch.Action
	setCurrentPeer *rematch.Action
}

func NewState() *AppState {
	// Generate the reducer from our model.
	Reducer := &AppModel{
		State: AppData{
			Peers:   make([]Peer, 0),
			Current: -1, // Should mean non selected.
		},
	}
	// Instantiate the redux state from the reducer.
	return &AppState{
		Reducer: Reducer,
		// Define the store.
		Store: store.New(Reducer),
		// Define available reducers for the store.
		setNodeInfo:    Reducer.Action(Reducer.SetInfo),
		updatePeers:    Reducer.Action(Reducer.Update),
		setCurrentPeer: Reducer.Action(Reducer.Current),
	}
}

// Helpers for shorter calls.
func (s *AppState) UpdateInfo(info NodeInfo) {
	s.Store.Dispatch(s.setNodeInfo.With(info))
}

func (s *AppState) UpdatePeers(peers []Peer) {
	s.Store.Dispatch(s.updatePeers.With(peers))
}

func (s *AppState) GetCurrent() *Peer {
	state := s.GetData()
	if state.Current == -1 {
		return nil
	}
	return &state.Peers[state.Current]
}

func (s *AppState) SetCurrentPeer(index int) {
	s.Store.Dispatch(s.setCurrentPeer.With(index))
}

func (s *AppState) GetData() AppData {
	return s.Store.StateOf(s.Reducer).(AppData)
}
