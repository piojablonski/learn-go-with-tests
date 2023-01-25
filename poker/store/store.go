package store

import "github.com/piojablonski/learn-go-with-tests/poker/business"

type PlayerStore interface {
	GetScoreByPlayer(name string) (score int, found bool)
	RecordWin(name string) error
	GetAllPlayers() business.League
}
