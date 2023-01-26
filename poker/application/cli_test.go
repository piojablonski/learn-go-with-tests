package application_test

import (
	"bytes"
	"fmt"
	"github.com/piojablonski/learn-go-with-tests/poker/application"
	"github.com/piojablonski/learn-go-with-tests/poker/common/testhelpers"

	"strings"
	"testing"
	"time"
)

var scores = map[string]int{}

var dummyStore = new(testhelpers.StubPlayerStore)
var dummyAlerter = new(SpyBlindAlerter)
var dummyIn = new(bytes.Buffer)
var dummyOut = new(bytes.Buffer)

func TestSaveScores(t *testing.T) {
	t.Run("registers Swiatek won", func(t *testing.T) {

		store := &testhelpers.StubPlayerStore{Scores: scores}

		in := strings.NewReader("Swiatek wins\n")

		cli := application.NewCLI(store, in, dummyOut, dummyAlerter)
		err := cli.PlayPoker()
		testhelpers.AssertNoError(t, err)

		testhelpers.AssertEqual(t, len(store.WinCalls), 1)

		want := "Swiatek"
		got := store.WinCalls[0]
		testhelpers.AssertEqual(t, got, want)
	})
	t.Run("asks for number of players and schedules proper amount of alerts", func(t *testing.T) {

		out := new(bytes.Buffer)
		in := bytes.NewBufferString("2\n")
		//in := strings.NewReader("2\n")
		fmt.Fprintln(in, "Swiatek wins")
		alerter := &SpyBlindAlerter{}
		cli := application.NewCLI(dummyStore, in, out, alerter)
		err := cli.PlayPoker()
		testhelpers.AssertNoError(t, err)

		gotInConsole := out.String()
		wantInConsole := application.PlayersPrompt
		testhelpers.AssertEqual(t, gotInConsole, wantInConsole)

		want := []application.ScheduleAlert{
			{0, 100},
			{7 * time.Minute, 200},
			{14 * time.Minute, 300},
		}

		got := alerter.alerts[:3]
		testhelpers.AssertEqual(t, got, want)

	})
}

type SpyBlindAlerter struct {
	alerts []application.ScheduleAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int) {
	s.alerts = append(s.alerts, application.ScheduleAlert{at, amount})
}

func TestBlindAlerts(t *testing.T) {
	store := &testhelpers.StubPlayerStore{Scores: scores}

	in := bytes.NewBufferString("5\n")
	fmt.Fprint(in, "Swiatek wins\n")

	blindAlerter := &SpyBlindAlerter{}
	cli := application.NewCLI(store, in, dummyOut, blindAlerter)
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
