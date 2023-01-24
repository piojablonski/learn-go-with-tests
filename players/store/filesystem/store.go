package filesystem

import (
	"encoding/json"
	"io"
	"os"
	"players/application"
	"players/business"
	"players/store"
)

type FilesystemStore struct {
	Database *json.Encoder
	league   business.League
}

func NewStore(db *os.File) store.PlayerStore {
	s := new(FilesystemStore)
	s.Database = json.NewEncoder(&store.Tape{File: db})
	db.Seek(0, io.SeekStart)
	players, _ := application.ReadPlayers(db)
	s.league = players
	return s
}

func (s *FilesystemStore) GetAllPlayers() business.League {
	return s.league
}

func (s *FilesystemStore) GetScoreByPlayer(name string) (int, bool) {
	player := s.league.Find(name)
	if player != nil {
		return player.Score, true
	} else {
		return 0, false
	}
}

func (s *FilesystemStore) RecordWin(name string) {
	player := s.league.Find(name)

	if player != nil {
		player.Score++
	} else {
		s.league = append(s.league, business.Player{Name: name, Score: 1})
	}
	s.Database.Encode(s.league)

}
