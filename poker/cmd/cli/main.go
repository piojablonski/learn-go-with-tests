package main

import (
	"fmt"
	"github.com/piojablonski/learn-go-with-tests/poker/application"
	"github.com/piojablonski/learn-go-with-tests/poker/store/filesystem"
	"log"
	"os"
	"time"
)

const dbFileName = "../db.json"

type DefaultBlindAlerter struct{}

func (d DefaultBlindAlerter) ScheduleAlertAt(_ time.Duration, _ int) {
	log.Println("implement me")
}

func main() {
	fmt.Println("poker command line utility")

	store, closeStore, err := filesystem.NewStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer closeStore()
	cli := application.NewCLI(store, os.Stdin, &DefaultBlindAlerter{})
	for {
		err := cli.PlayPoker()
		if err != nil {
			log.Fatal(err)
		}
	}
}
