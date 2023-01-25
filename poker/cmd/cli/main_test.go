package main_test

import (
	main "github.com/piojablonski/learn-go-with-tests/poker/cmd/cli"
	"github.com/piojablonski/learn-go-with-tests/poker/common/testhelpers"
	"strings"
	"testing"
)

var scores = map[string]int{
	//"Swiatek": 300,
	//"Hurkacz": 234,
	//"Kubot":   0,
}

func TestSaveScores(t *testing.T) {
	store := &testhelpers.StubPlayerStore{Scores: scores}

	in := strings.NewReader("Swiatek wins\n")

	cli := main.NewCLI(store, in)
	cli.PlayPoker()

	testhelpers.AssertEqual(t, len(store.WinCalls), 1)

	want := "Swiatek"
	got := store.WinCalls[0]
	testhelpers.AssertEqual(t, got, want)
}
