package store

import "players/business"

type PlayerStore interface {
	GetScoreByPlayer(name string) (score int, found bool)
	RecordWin(name string) error
	GetAllPlayers() business.League
}
