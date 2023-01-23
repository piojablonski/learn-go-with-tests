package store

type PlayerStore interface {
	GetScoreByPlayer(name string) (score int, found bool)
	RecordWin(name string)
}
