package store

import (
	"encoding/json"
	"errors"
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
	if player != nil {
		return player.Score, nil
	} else {
		return 0, errors.New("player not found")
	}
}

func (s *FilesystemStore) RecordWin(name string) {
	players := s.GetAllPlayers()
	player := players.Find(name)

	if player != nil {
		player.Score++
	} else {
		players = append(players, business.Player{Name: name, Score: 1})
	}
	s.Database.Seek(0, io.SeekStart)
	json.NewEncoder(s.Database).Encode(players)

}
