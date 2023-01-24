package main

import (
	"log"
	"net/http"
	"os"
	"players/httpserver"
	"players/store/filesystem"
)

func main() {
	f, err := os.OpenFile("db", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("cannot open file %q", err)
	}
	store := filesystem.NewStore(f)
	err = http.ListenAndServe(":8080", httpserver.NewPlayerServer(store))
	if err != nil {
		log.Fatal(err)
	}

}
