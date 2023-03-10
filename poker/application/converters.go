package application

import (
	"encoding/json"
	"fmt"
	"github.com/piojablonski/learn-go-with-tests/poker/business"
	"io"
)

func ReadPlayers(r io.Reader) ([]business.Player, error) {
	var players []business.Player

	// json decodes json string to golang player type
	err := json.NewDecoder(r).Decode(&players)

	if err != nil {
		err = fmt.Errorf("parsing players failed, %w", err)
	}

	return players, err

}
