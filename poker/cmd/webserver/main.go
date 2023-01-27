package main

import (
	"fmt"
	"github.com/piojablonski/learn-go-with-tests/poker/application"
	"github.com/piojablonski/learn-go-with-tests/poker/httpserver"
	"github.com/piojablonski/learn-go-with-tests/poker/store/filesystem"
	"log"
	"net/http"
	"time"
)

func alerter(at time.Duration, amount int) {
	time.AfterFunc(at, func() {
		fmt.Println("ws", amount)
	})
}

func main() {
	const dbFileName = "../db.json"
	store, close, err := filesystem.NewStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	game := application.NewGame(store, application.BlindAlerterFunc(alerter))
	defer close()
	srv, err := httpserver.NewPlayerServer(store, game)
	if err != nil {
		log.Fatal(err)
	}
	err = http.ListenAndServe(":8080", srv)
	if err != nil {
		log.Fatal(err)
	}

}
