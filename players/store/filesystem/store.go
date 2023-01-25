package filesystem

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"players/application"
	"players/business"
	"players/store"
	"sort"
)

type FilesystemStore struct {
	Database *json.Encoder
	league   business.League
}

func NewStore(file *os.File) (store.PlayerStore, error) {
	s := new(FilesystemStore)
	s.Database = json.NewEncoder(&store.Tape{File: file})
	_, err := file.Seek(0, io.SeekStart)
	if err != nil {
		return nil, fmt.Errorf("problem seeking in file %w", err)
	}

	info, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("problem getting info for file %s, %w", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}

	players, err := application.ReadPlayers(file)
	if err != nil {
		return nil, fmt.Errorf("problem reading players from file, %w", err)
	}
	s.league = players
	return s, nil
}

func (s *FilesystemStore) GetAllPlayers() business.League {
	sort.Slice(s.league, func(i, j int) bool {
		if s.league[i].Score > s.league[j].Score {
			return true
		}
		return false
	})
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

func (s *FilesystemStore) RecordWin(name string) error {
	player := s.league.Find(name)

	if player != nil {
		player.Score++
	} else {
		s.league = append(s.league, business.Player{Name: name, Score: 1})
	}
	err := s.Database.Encode(s.league)
	if err != nil {
		return fmt.Errorf("problem recording win for a player, %w", err)
	}
	return nil

}
