package application_test

import (
	"fmt"
	"github.com/piojablonski/learn-go-with-tests/poker/application"
	"github.com/piojablonski/learn-go-with-tests/poker/common/testhelpers"
	"io"
	"testing"
	"time"
)

func TestGame(t *testing.T) {

	t.Run("schedules proper alerts for 2 players", func(t *testing.T) {
		alerter := &SpyBlindAlerter{}
		game := application.NewGame(dummyStore, alerter)
		game.StartGame(2, io.Discard)

		want := []application.ScheduledAlert{
			{0, 100},
			{7 * time.Minute, 200},
			{14 * time.Minute, 300},
		}

		got := alerter.alerts[:3]
		testhelpers.AssertEqual(t, got, want)

	})
	t.Run("schedules all alerts for 5 players", func(t *testing.T) {
		store := &testhelpers.SpyPlayerStore{Scores: scores}
		blindAlerter := &SpyBlindAlerter{}
		game := application.NewGame(store, blindAlerter)
		game.StartGame(5, io.Discard)

		cases := []application.ScheduledAlert{
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
	})

	t.Run("finishing game should store a winner", func(t *testing.T) {
		store := &testhelpers.SpyPlayerStore{Scores: scores}
		blindAlerter := &SpyBlindAlerter{}
		game := application.NewGame(store, blindAlerter)
		game.StartGame(5, io.Discard)
		game.Finish("Piotr")

		got := store.WinCalls[0]
		want := "Piotr"

		testhelpers.AssertEqual(t, got, want)

	})
}
