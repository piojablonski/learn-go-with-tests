package application_test

import (
	"bytes"
	"github.com/piojablonski/learn-go-with-tests/poker/application"
	"github.com/piojablonski/learn-go-with-tests/poker/common/testhelpers"
	"io"

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

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int, _ io.Writer) {
	s.alerts = append(s.alerts, application.ScheduledAlert{at, amount})
}

func TestSaveScores(t *testing.T) {
	t.Run("start with 2 players and finish with Swiatek as a winner", func(t *testing.T) {
		in := strings.NewReader("2\nSwiatek wins\n")
		game := &application.SpyGame{}
		cli := application.NewCLI(in, dummyOut, game)
		err := cli.PlayPoker()
		testhelpers.AssertNoError(t, err)
		application.AssertReceivedMessages(t, dummyOut, application.PlayersPrompt)
		application.AssertFinishCalledWith(t, game, "Swiatek")
		application.AssertGameStartedWithPlayers(t, game, 2)
	})
	t.Run("start with 8 players and finish with Zimoch as a winner", func(t *testing.T) {
		in := strings.NewReader("8\nZimoch wins\n")
		game := &application.SpyGame{}
		cli := application.NewCLI(in, dummyOut, game)
		err := cli.PlayPoker()
		testhelpers.AssertNoError(t, err)
		application.AssertReceivedMessages(t, dummyOut, application.PlayersPrompt)
		application.AssertFinishCalledWith(t, game, "Zimoch")
		application.AssertGameStartedWithPlayers(t, game, 8)
	})

	t.Run("handles inproper format of number of players", func(t *testing.T) {
		in := strings.NewReader("five\n")
		out := new(bytes.Buffer)
		game := new(application.SpyGame)
		err := application.NewCLI(in, out, game).PlayPoker()
		testhelpers.AssertNoError(t, err)

		application.AssertGameNotStarted(t, game)
		application.AssertReceivedMessages(t, out, application.PlayersPrompt, application.WrongNumberOfPlayers)

	})

	t.Run("displays an error when an incorrect text is typed to finish the game", func(t *testing.T) {
		in := strings.NewReader("2\nSarah is a killer")
		out := new(bytes.Buffer)
		game := new(application.SpyGame)
		err := application.NewCLI(in, out, game).PlayPoker()
		testhelpers.AssertNoError(t, err)
		application.AssertGameStartedWithPlayers(t, game, 2)
		application.AssertReceivedMessages(t, out, application.PlayersPrompt, application.ErrFinishingGameWrongInput)
		application.AssertGameNotFinished(t, game)
	})
}
