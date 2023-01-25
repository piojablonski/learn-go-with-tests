package main_test

import (
	"github.com/piojablonski/learn-go-with-tests/poker/common/testhelpers"
	"testing"
)
import cli "github.com/piojablonski/learn-go-with-tests/poker/cmd/cli"

var scores = map[string]int{
	"Swiatek": 300,
	"Hurkacz": 234,
	"Kubot":   0,
}

func TestSaveScores(t *testing.T) {
	store := &testhelpers.StubPlayerStore{Scores: scores}

	cli.RecordWin(store, "Piotr")
	testhelpers.AssertEqual(t, len(store.WinCalls), 1)
}
