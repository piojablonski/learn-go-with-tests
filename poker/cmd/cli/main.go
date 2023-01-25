package main

import (
	"fmt"
	"github.com/piojablonski/learn-go-with-tests/poker/store"
	"github.com/piojablonski/learn-go-with-tests/poker/store/filesystem"
	"io"
	"log"
	"os"
	"time"
)

const dbFileName = "../db.json"

func main() {
	fmt.Println("poker command line utility")

	store, close, err := filesystem.NewStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()
	cli := NewCLI(store, os.Stdin, &DefaultBlindAlerter{})
	for {
		cli.PlayPoker()
	}
}

func (cli *CLI) PlayPoker() {
	var name string
	fmt.Fscanln(cli.in, &name)
	cli.alerter.ScheduleAlertAt(5*time.Minute, 100)
	cli.store.RecordWin(name)
}

func NewCLI(store store.PlayerStore, in io.Reader, alerter BlindAlerter) *CLI {
	return &CLI{store, in, alerter}
}

type CLI struct {
	store   store.PlayerStore
	in      io.Reader
	alerter BlindAlerter
}

type BlindAlerter interface {
	ScheduleAlertAt(at time.Duration, amount int)
}

type DefaultBlindAlerter struct{}

func (d DefaultBlindAlerter) ScheduleAlertAt(at time.Duration, amount int) {
	log.Println("implement me")
}
