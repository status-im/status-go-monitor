package internal

import (
	"github.com/dannypsnl/redux/v2/rematch"
	"github.com/dannypsnl/redux/v2/store"
)

// AppState is a wrapper around the Redux store
type AppState struct {
	Reducer        *AppModel
	Store          *store.Store
	setNodeInfo    *rematch.Action
	updatePeers    *rematch.Action
	setCurrentPeer *rematch.Action
}

// NewState creates a new instance of the redux store
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
