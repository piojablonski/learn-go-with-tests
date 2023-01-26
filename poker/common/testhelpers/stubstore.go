package testhelpers

import "github.com/piojablonski/learn-go-with-tests/poker/business"

type SpyPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
	League   []business.Player
}

func (s *SpyPlayerStore) GetScoreByPlayer(name string) (score int, found bool) {
	score, found = s.Scores[name]
	return
}

func (s *SpyPlayerStore) RecordWin(name string) error {
	s.WinCalls = append(s.WinCalls, name)
	return nil
}

func (s *SpyPlayerStore) GetAllPlayers() business.League {
	return s.League
}
