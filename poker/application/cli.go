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

var Amounts = []int{100, 200,
	300,
	400,
	500,
	600,
	800,
	1000,
	2000,
	4000,
	8000}

func (s *ScheduleAlert) String() string {
	return fmt.Sprintf("%d chips at %s", s.Amount, s.ScheduledAt)
}

func NewCLI(store store.PlayerStore, in io.Reader, out io.Writer, alerter BlindAlerter) *CLI {
	return &CLI{store, in, alerter, out}
}

type CLI struct {
	store   store.PlayerStore
	in      io.Reader
	alerter BlindAlerter
	out     io.Writer
}

const PlayersPrompt = "Podaj ilość graczy: "

func (cli *CLI) PlayPoker() error {
	var (
		name        string
		noOfPlayers int
	)

	fmt.Fprint(cli.out, PlayersPrompt)
	fmt.Fscanln(cli.in, &noOfPlayers)
	//if err != nil {
	//	return fmt.Errorf("playing poker, problem while scanning number of players, %w", err)
	//}

	cli.scheduleAllAlerts(noOfPlayers)
	fmt.Fscanln(cli.in, &name)
	//if err != nil {
	//	return fmt.Errorf("playing poker, problem while scanning name, %w", err)
	//}
	err := cli.store.RecordWin(name)
	if err != nil {
		return fmt.Errorf("playing poker, problem recording name, %w", err)
	}
	return nil
}

func (cli *CLI) scheduleAllAlerts(noOfPlayers int) {
	timeIncrement := time.Duration(5+noOfPlayers) * time.Minute
	blindTime := 0 * time.Second
	for _, amount := range Amounts {
		//ti := timeIncrement*i + 10*time.Secon
		cli.alerter.ScheduleAlertAt(blindTime, amount)
		blindTime = blindTime + timeIncrement
	}
}

type BlindAlerter interface {
	ScheduleAlertAt(at time.Duration, amount int)
}

type BlindAlerterFunc func(at time.Duration, amount int)

func (b BlindAlerterFunc) ScheduleAlertAt(at time.Duration, amount int) {
	b(at, amount)
}
