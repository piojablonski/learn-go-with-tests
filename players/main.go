package main

import (
	"log"
	"net/http"
	"os"
	"players/httpserver"
	"players/store/filesystem"
)

func main() {
	f, err := os.Create("db")
	if err != nil {
		log.Fatalf("cannot open file %q", err)
	}
	store := filesystem.NewStore(f)
	err = http.ListenAndServe(":8080", httpserver.NewPlayerServer(store))
	if err != nil {
		log.Fatal(err)
	}

}
