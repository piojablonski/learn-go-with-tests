package main

import (
	"github.com/piojablonski/learn-go-with-tests/poker/httpserver"
	"github.com/piojablonski/learn-go-with-tests/poker/store/filesystem"
	"log"
	"net/http"
)

func main() {
	const dbFileName = "../db.json"
	store, close, err := filesystem.NewStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()
	srv, err := httpserver.NewPlayerServer(store)
	if err != nil {
		log.Fatal(err)
	}
	err = http.ListenAndServe(":8080", srv)
	if err != nil {
		log.Fatal(err)
	}

}
