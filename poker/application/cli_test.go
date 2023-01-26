package application_test

import (
	"fmt"
	"github.com/piojablonski/learn-go-with-tests/poker/application"
	"github.com/piojablonski/learn-go-with-tests/poker/common/testhelpers"

	"strings"
	"testing"
	"time"
)

var scores = map[string]int{}

func TestSaveScores(t *testing.T) {
	store := &testhelpers.StubPlayerStore{Scores: scores}
	dummyBlindAlerter := &SpyBlindAlerter{}

	in := strings.NewReader("Swiatek wins\n")

	cli := application.NewCLI(store, in, dummyBlindAlerter)
	err := cli.PlayPoker()
	testhelpers.AssertNoError(t, err)

	testhelpers.AssertEqual(t, len(store.WinCalls), 1)

	want := "Swiatek"
	got := store.WinCalls[0]
	testhelpers.AssertEqual(t, got, want)
}

type SpyBlindAlerter struct {
	alerts []application.ScheduleAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int) {
	s.alerts = append(s.alerts, application.ScheduleAlert{at, amount})
}

func TestBlindAlerts(t *testing.T) {
	store := &testhelpers.StubPlayerStore{Scores: scores}

	in := strings.NewReader("Swiatek wins\n")

	blindAlerter := &SpyBlindAlerter{}
	cli := application.NewCLI(store, in, blindAlerter)
	err := cli.PlayPoker()
	testhelpers.AssertNoError(t, err)

	cases := []application.ScheduleAlert{
		{0 * time.Second, 100},
		{10 * time.Minute, 200},
		{20 * time.Minute, 300},
		{30 * time.Minute, 400},
		{40 * time.Minute, 500},
		{50 * time.Minute, 600},
		{60 * time.Minute, 800},
		{70 * time.Minute, 1000},
		{80 * time.Minute, 2000},
		{90 * time.Minute, 4000},
		{100 * time.Minute, 8000},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d scheduled for %v", c.Amount, c.ScheduledAt), func(t *testing.T) {

			if len(blindAlerter.alerts) <= i {
				t.Fatalf("alert %d has not been set for %v", i, c)
			}

			got := blindAlerter.alerts[i]
			testhelpers.AssertEqual(t, got.ScheduledAt, c.ScheduledAt)
			testhelpers.AssertEqual(t, got.Amount, c.Amount)
		})
	}
}