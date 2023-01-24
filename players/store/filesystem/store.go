package store

import (
	"encoding/json"
	"io"
	"players/application"
	"players/business"
)

type FilesystemStore struct {
	Database io.ReadWriteSeeker
}

func (s *FilesystemStore) GetAllPlayers() business.League {
	s.Database.Seek(0, io.SeekStart)
	players, _ := application.ReadPlayers(s.Database)
	return players
}

func (s *FilesystemStore) GetScoreByPlayer(name string) (int, error) {
	player := s.GetAllPlayers().Find(name)
	return player.Score, nil
}

func (s *FilesystemStore) RecordWin(name string) {
	players := s.GetAllPlayers()
	player := players.Find(name)
	player.Score++
	s.Database.Seek(0, io.SeekStart)
	json.NewEncoder(s.Database).Encode(players)

}
