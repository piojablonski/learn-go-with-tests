package application

import (
	"fmt"
	"github.com/piojablonski/learn-go-with-tests/poker/store"
	"io"
	"time"
)

type ScheduleAlert struct {
	ScheduledAt time.Duration
	Amount      int
}

var Schedule = []ScheduleAlert{{0 * time.Second, 100},
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

func (s *ScheduleAlert) String() string {
	return fmt.Sprintf("%d chips at %s", s.Amount, s.ScheduledAt)
}

func NewCLI(store store.PlayerStore, in io.Reader, alerter BlindAlerter) *CLI {
	return &CLI{store, in, alerter}
}

type CLI struct {
	store   store.PlayerStore
	in      io.Reader
	alerter BlindAlerter
}

func (cli *CLI) PlayPoker() error {
	var name string
	_, err := fmt.Fscan(cli.in, &name)
	if err != nil {
		return fmt.Errorf("playing poker, problem while scanning name, %w", err)
	}
	cli.scheduleAllAlerts()
	err = cli.store.RecordWin(name)
	if err != nil {
		return fmt.Errorf("playing poker, problem recording name, %w", err)
	}
	return nil
}

func (cli *CLI) scheduleAllAlerts() {
	for _, s := range Schedule {
		cli.alerter.ScheduleAlertAt(s.ScheduledAt, s.Amount)
	}
}

type BlindAlerter interface {
	ScheduleAlertAt(at time.Duration, amount int)
}
