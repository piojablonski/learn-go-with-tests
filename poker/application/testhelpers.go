package application

import (
	"bytes"
	"github.com/piojablonski/learn-go-with-tests/poker/common/testhelpers"
	"strings"
	"testing"
)

type SpyGame struct {
	CallsToStartGame  []int
	CallsToFinishGame []string
}

func (s *SpyGame) StartGame(noOfPlayers int) {
	s.CallsToStartGame = append(s.CallsToStartGame, noOfPlayers)
}

func (s *SpyGame) Finish(winner string) error {
	s.CallsToFinishGame = append(s.CallsToFinishGame, winner)
	return nil
}

func AssertGameNotStarted(t *testing.T, game *SpyGame) {
	if len(game.CallsToStartGame) > 0 {
		t.Errorf("game should not have started")
	}
}

func AssertGameStartedWithPlayers(t *testing.T, game *SpyGame, noOfPlayers int) {
	got2 := game.CallsToStartGame
	want2 := []int{noOfPlayers}
	testhelpers.AssertEqual(t, got2, want2)
}

func AssertFinishCalledWith(t *testing.T, game *SpyGame, name string) {
	want := []string{name}
	got := game.CallsToFinishGame
	testhelpers.AssertEqual(t, got, want)
}

func AssertGameNotFinished(t *testing.T, game *SpyGame) {
	if len(game.CallsToFinishGame) > 0 {
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
