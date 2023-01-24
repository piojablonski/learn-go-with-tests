package inmemory

import (
	. "players/business"
	"players/store"
)

// var scores = map[string]int{
// 	// "Radwańska": 300,
// 	// "Hurkacz":   234,
// }

// implements PlayerStore
type InmemoryPlayerStore struct {
	scores map[string]int
}

func NewInmemoryPlayerStore() store.PlayerStore {
	var scores = map[string]int{}
	return &InmemoryPlayerStore{scores}
}

func (s *InmemoryPlayerStore) GetScoreByPlayer(name string) (score int, found bool) {
	score, found = s.scores[name]
	return
}
func (s *InmemoryPlayerStore) RecordWin(name string) {
	s.scores[name]++
}

func (s *InmemoryPlayerStore) GetAllPlayers() []Player {

	var players []Player
	for name, score := range s.scores {
		players = append(players, Player{Name: name, Score: score})
	}
	return players
}

// func (s *InmemoryPlayerStore) GetAll() []Player {

// }
