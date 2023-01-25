package httpserver

import (
	"encoding/json"
	"fmt"
	"github.com/piojablonski/learn-go-with-tests/poker/business"
	"github.com/piojablonski/learn-go-with-tests/poker/store"
	"net/http"
	"strings"
)

type PlayerServer struct {
	store store.PlayerStore
	// it includes serve http to the struct on the root level
	http.Handler
}

func NewPlayerServer(store store.PlayerStore) *PlayerServer {
	srv := new(PlayerServer)
	router := http.NewServeMux()
	// router and PlayerServer both implements Handle interface so I can assign the router to the server
	srv.Handler = router

	router.HandleFunc("/league", srv.leagueHandler)
	router.HandleFunc("/player/", srv.playerHandler)
	srv.store = store
	return srv

}

const ApplicationJsonContentType = "application/json"

func (ps *PlayerServer) leagueHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", ApplicationJsonContentType)
	json.NewEncoder(w).Encode(ps.getLeagueTable())
}

func (ps *PlayerServer) getLeagueTable() []business.Player {
	players := ps.store.GetAllPlayers()
	return players

}

func (ps *PlayerServer) playerHandler(w http.ResponseWriter, req *http.Request) {

	switch req.Method {

	case http.MethodPost:
		ps.saveScores(w, req)

	case http.MethodGet:
		ps.showScores(w, req)
	}
}

func (ps *PlayerServer) saveScores(res http.ResponseWriter, req *http.Request) {

	name := getPlayerName(req)

	err := ps.store.RecordWin(name)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}
}

func getPlayerName(req *http.Request) string {
	paths := strings.Split(req.URL.Path, "/")
	name := paths[len(paths)-1]
	return name
}

func (ps *PlayerServer) showScores(res http.ResponseWriter, req *http.Request) {
	name := getPlayerName(req)
	score, found := ps.store.GetScoreByPlayer(name)
	if !found {
		res.WriteHeader(http.StatusNotFound)
	} else {
		fmt.Fprintf(res, "%d", score)
	}
}
