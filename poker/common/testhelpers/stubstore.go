package testhelpers

import "github.com/piojablonski/learn-go-with-tests/poker/business"

type StubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
	League   []business.Player
}

func (s *StubPlayerStore) GetScoreByPlayer(name string) (score int, found bool) {
	score, found = s.Scores[name]
	return
}

func (s *StubPlayerStore) RecordWin(name string) error {
	s.WinCalls = append(s.WinCalls, "win")
	return nil
}

func (s *StubPlayerStore) GetAllPlayers() business.League {
	return s.League
}
