package main

import (
	"log"
	"net/http"
	"players/httpserver"
	"players/store/inmemory"
)

func main() {
	// http.HandleFunc("/player", httpserver.HandleGetPlayer)
	err := http.ListenAndServe(":8080", httpserver.NewPlayerServer(inmemory.NewInmemoryPlayerStore()))
	if err != nil {
		log.Fatal(err)
	}

}
