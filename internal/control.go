package internal

import (
	"log"
)

// StateController is the central point for state control
type StateController struct {
	State  *AppState
	Client *StatusGoClient
}

// Fetch for getting current state of peers from status-go
func (s *StateController) Fetch() {
	peers, err := s.Client.getPeers()
	if err != nil {
		log.Panicln(err)
	}
	ps := s.GetData()
	s.UpdatePeers(peers)
	if ps.Current == -1 {
		s.SetCurrentPeer(0)
	}
}

// TrustPeer marking the selected whisper peer as trusted
func (s *StateController) TrustPeer(peer *Peer) error {
	success, err := s.Client.trustPeer(peer.Enode)
	if err != nil {
		log.Panicln(err)
	}
	if success != true {
		log.Panicln("failed to trust whisper peer")
	}
	return nil
}

// RemovePeer for removing a selected peer from connected to status-go
func (s *StateController) RemovePeer(peer *Peer) error {
	success, err := s.Client.removePeer(peer.Enode)
	if err != nil || success != true {
		log.Panicln("failed to remove peer:", err)
		return err
	}
	s.Fetch()
	return nil
}

// GetInfo for getting information about current connected to node
func (s *StateController) GetInfo() error {
	info, err := s.Client.nodeInfo()
	if err != nil {
		return err
	}
	s.UpdateInfo(info)
	return nil
}

// Helpers for shorter calls.
func (s *StateController) UpdateInfo(info NodeInfo) {
	s.State.Store.Dispatch(s.State.setNodeInfo.With(info))
}

func (s *StateController) UpdatePeers(peers []Peer) {
	s.State.Store.Dispatch(s.State.updatePeers.With(peers))
}

func (s *StateController) GetCurrent() *Peer {
	state := s.GetData()
	if state.Current == -1 {
		return nil
	}
	return &state.Peers[state.Current]
}

func (s *StateController) SetCurrentPeer(index int) {
	s.State.Store.Dispatch(s.State.setCurrentPeer.With(index))
}

func (s *StateController) GetData() AppData {
	return s.State.Store.StateOf(s.State.Reducer).(AppData)
}
