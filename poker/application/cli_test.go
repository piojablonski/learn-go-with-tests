package application_test

import (
	"bytes"
	"github.com/piojablonski/learn-go-with-tests/poker/application"
	"github.com/piojablonski/learn-go-with-tests/poker/common/testhelpers"

	"strings"
	"testing"
	"time"
)

var scores = map[string]int{}

var dummyStore = new(testhelpers.SpyPlayerStore)
var dummyAlerter = new(SpyBlindAlerter)
var dummyIn = new(bytes.Buffer)
var dummyOut = new(bytes.Buffer)

type SpyBlindAlerter struct {
	alerts []application.ScheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int) {
	s.alerts = append(s.alerts, application.ScheduledAlert{at, amount})
}

type SpyGame struct {
	callsToStartGame  []int
	callsToFinishGame []string
}

func (s *SpyGame) StartGame(noOfPlayers int) {
	s.callsToStartGame = append(s.callsToStartGame, noOfPlayers)
}

func (s *SpyGame) Finish(winner string) error {
	s.callsToFinishGame = append(s.callsToFinishGame, winner)
	return nil
}

func TestSaveScores(t *testing.T) {
	t.Run("registers Swiatek won", func(t *testing.T) {
		in := strings.NewReader("Swiatek wins\n")
		game := &SpyGame{}
		cli := application.NewCLI(in, dummyOut, game)
		err := cli.PlayPoker()
		testhelpers.AssertNoError(t, err)
		want := []string{"Swiatek"}
		got := game.callsToFinishGame
		testhelpers.AssertEqual(t, got, want)
	})
	t.Run("reads amount of players and starts a game", func(t *testing.T) {
		out := new(bytes.Buffer)
		in := bytes.NewBufferString("2\n")
		game := new(SpyGame)
		cli := application.NewCLI(in, out, game)
		err := cli.PlayPoker()
		testhelpers.AssertNoError(t, err)

		got := game.callsToStartGame
		want := []int{2}
		testhelpers.AssertEqual(t, got, want)

	})

}
