package main_test

import (
	main "github.com/piojablonski/learn-go-with-tests/poker/cmd/cli"
	"github.com/piojablonski/learn-go-with-tests/poker/common/testhelpers"
	"strings"
	"testing"
	"time"
)

var scores = map[string]int{
	//"Swiatek": 300,
	//"Hurkacz": 234,
	//"Kubot":   0,
}

func TestSaveScores(t *testing.T) {
	store := &testhelpers.StubPlayerStore{Scores: scores}
	dummyBlindAlerter := &SpyBlindAlerter{}

	in := strings.NewReader("Swiatek wins\n")

	cli := main.NewCLI(store, in, dummyBlindAlerter)
	cli.PlayPoker()

	testhelpers.AssertEqual(t, len(store.WinCalls), 1)

	want := "Swiatek"
	got := store.WinCalls[0]
	testhelpers.AssertEqual(t, got, want)
}

type SpyBlindAlerter struct {
	alerts []struct {
		scheduledAt time.Duration
		amount      int
	}
}

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int) {
	s.alerts = append(s.alerts, struct {
		scheduledAt time.Duration
		amount      int
	}{at, amount})
}

func TestBlindAlerts(t *testing.T) {
	store := &testhelpers.StubPlayerStore{Scores: scores}

	in := strings.NewReader("Swiatek wins\n")

	blindAlerter := &SpyBlindAlerter{}
	cli := main.NewCLI(store, in, blindAlerter)
	cli.PlayPoker()

	got := len(blindAlerter.alerts)
	if got != 1 {
		t.Fatalf("expected to receive 1 alert, but got %d", got)
	}
}
