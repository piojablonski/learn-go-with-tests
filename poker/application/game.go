package application

import (
	"fmt"
	"github.com/piojablonski/learn-go-with-tests/poker/store"
	"io"
	"time"
)

type Game interface {
	StartGame(noOfPlayers int, w io.Writer)
	Finish(winner string) error
}

type DefaultGame struct {
	store   store.PlayerStore
	alerter BlindAlerter
}

func (g *DefaultGame) StartGame(noOfPlayers int, w io.Writer) {
	g.scheduleAllAlerts(noOfPlayers, w)
}

func (g *DefaultGame) Finish(winner string) error {
	err := g.store.RecordWin(winner)
	if err != nil {
		return fmt.Errorf("problem recording a winner, %w", err)
	}
	return nil
}

func NewGame(playerStore store.PlayerStore, alerter BlindAlerter) *DefaultGame {
	return &DefaultGame{playerStore, alerter}

}

var blindAmounts = []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}

func (g *DefaultGame) scheduleAllAlerts(noOfPlayers int, w io.Writer) {
	timeIncrement := time.Duration(5+noOfPlayers) * time.Minute
	blindTime := 0 * time.Second
	for _, amount := range blindAmounts {
		g.alerter.ScheduleAlertAt(blindTime, amount, w)
		blindTime = blindTime + timeIncrement
	}
}
