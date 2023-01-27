package main

import (
	"fmt"
	"github.com/piojablonski/learn-go-with-tests/poker/application"
	"github.com/piojablonski/learn-go-with-tests/poker/store/filesystem"
	"io"
	"log"
	"os"
	"time"
)

const dbFileName = "../db.json"

func StdOutBlindAlerter(at time.Duration, amount int, w io.Writer) {
	time.AfterFunc(at, func() {
		_, err := fmt.Fprintf(w, "current blind is: %d", amount)
		if err != nil {
			log.Fatal(err)
		}
	})
}

func main() {
	fmt.Println("You are playing poker!")

	store, closeStore, err := filesystem.NewStoreFromFile(dbFileName)
	game := application.NewGame(store, application.BlindAlerterFunc(StdOutBlindAlerter))
	if err != nil {
		log.Fatal(err)
	}
	defer closeStore()
	cli := application.NewCLI(os.Stdin, os.Stdout, game)
	for {
		err := cli.PlayPoker()
		if err != nil {
			log.Fatal(err)
		}
	}
}
