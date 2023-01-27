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
	t.Run("start with 2 players and finish with Swiatek as a winner", func(t *testing.T) {
		in := strings.NewReader("2\nSwiatek wins\n")
		game := &SpyGame{}
		cli := application.NewCLI(in, dummyOut, game)
		err := cli.PlayPoker()
		testhelpers.AssertNoError(t, err)
		AssertReceivedMessages(t, dummyOut, application.PlayersPrompt)
		AssertFinishCalledWith(t, game, "Swiatek")
		AssertGameStartedWithPlayers(t, game, 2)
	})
	t.Run("start with 8 players and finish with Zimoch as a winner", func(t *testing.T) {
		in := strings.NewReader("8\nZimoch wins\n")
		game := &SpyGame{}
		cli := application.NewCLI(in, dummyOut, game)
		err := cli.PlayPoker()
		testhelpers.AssertNoError(t, err)
		AssertReceivedMessages(t, dummyOut, application.PlayersPrompt)
		AssertFinishCalledWith(t, game, "Zimoch")
		AssertGameStartedWithPlayers(t, game, 8)
	})

	t.Run("handles inproper format of number of players", func(t *testing.T) {
		in := strings.NewReader("five\n")
		out := new(bytes.Buffer)
		game := new(SpyGame)
		err := application.NewCLI(in, out, game).PlayPoker()
		testhelpers.AssertNoError(t, err)

		AssertGameNotStarted(t, game)
		AssertReceivedMessages(t, out, application.PlayersPrompt, application.WrongNumberOfPlayers)

	})

	t.Run("displays an error when an incorrect text is typed to finish the game", func(t *testing.T) {
		in := strings.NewReader("2\nSarah is a killer")
		out := new(bytes.Buffer)
		game := new(SpyGame)
		err := application.NewCLI(in, out, game).PlayPoker()
		testhelpers.AssertNoError(t, err)
		AssertGameStartedWithPlayers(t, game, 2)
		AssertReceivedMessages(t, out, application.PlayersPrompt, application.ErrFinishingGameWrongInput)
		AssertGameNotFinished(t, game)
	})
}

func AssertGameNotStarted(t *testing.T, game *SpyGame) {
	if len(game.callsToStartGame) > 0 {
		t.Errorf("game should not have started")
	}
}

func AssertGameStartedWithPlayers(t *testing.T, game *SpyGame, noOfPlayers int) {
	got2 := game.callsToStartGame
	want2 := []int{noOfPlayers}
	testhelpers.AssertEqual(t, got2, want2)
}

func AssertFinishCalledWith(t *testing.T, game *SpyGame, name string) {
	want := []string{name}
	got := game.callsToFinishGame
	testhelpers.AssertEqual(t, got, want)
}

func AssertGameNotFinished(t *testing.T, game *SpyGame) {
	if len(game.callsToFinishGame) > 0 {
		t.Fatalf("Game shouldn't be finished")
	}
}

func AssertReceivedMessages(t *testing.T, out *bytes.Buffer, messages ...string) {

	want := strings.Join(messages, "\n")
	got := out.String()
	if !strings.Contains(got, want) {
		t.Errorf("Expected to see %q", want)
	}

}
