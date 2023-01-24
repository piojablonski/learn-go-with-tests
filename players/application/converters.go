package application

import (
	"encoding/json"
	"fmt"
	"io"
	"players/business"
)

func ReadPlayers(r io.Reader) ([]business.Player, error) {
	var players []business.Player

	// json decodes json string to golang player type
	err := json.NewDecoder(r).Decode(&players)

	if err != nil {
		err = fmt.Errorf("problem parsing players: %q", err)
	}

	return players, err

}
