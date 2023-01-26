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

func StdOutBlindAlerter(at time.Duration, amount int) {
	time.AfterFunc(at, func() {
		_, err := fmt.Fprintf(os.Stdout, "current blind is: %d", amount)
		if err != nil {
			log.Fatal(err)
		}
	})
}

func main() {
	fmt.Println("You are playing poker!")

	store, closeStore, err := filesystem.NewStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer closeStore()
	cli := application.NewCLI(store, os.Stdin, application.BlindAlerterFunc(StdOutBlindAlerter))
	for {
		err := cli.PlayPoker()
		if err != nil {
			log.Fatal(err)
		}
	}
}
