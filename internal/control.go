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
	ps := s.State.GetData()
	s.State.UpdatePeers(peers)
	if ps.Current == -1 {
		s.State.SetCurrentPeer(0)
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
	if err != nil {
		log.Panicln(err)
	}
	if success != true {
		log.Panicln("failed to remove peer")
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
	s.State.UpdateInfo(info)
	return nil
}
