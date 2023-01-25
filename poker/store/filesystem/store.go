package filesystem

import (
	"encoding/json"
	"fmt"
	"github.com/piojablonski/learn-go-with-tests/poker/application"
	"github.com/piojablonski/learn-go-with-tests/poker/business"
	"github.com/piojablonski/learn-go-with-tests/poker/store"
	"io"
	"os"
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

func NewStoreFromFile(name string) (store.PlayerStore, func(), error) {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot open file while creating store, %w", err)
	}
	store, err := NewStore(f)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot create store %w", err)
	}
	close := func() {
		f.Close()
	}
	return store, close, nil
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
