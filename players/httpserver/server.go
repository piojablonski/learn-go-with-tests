package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"players/business"
	"players/store"
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

func (ps *PlayerServer) leagueHandler(w http.ResponseWriter, req *http.Request) {

	json.NewEncoder(w).Encode(ps.getLeagueTable())
	w.WriteHeader(http.StatusOK)
}

func (ps *PlayerServer) getLeagueTable() []business.Player {
	players := []business.Player{
		{Name: "Swiatek", Score: 300},
		{Name: "Hurkacz", Score: 234},
		{Name: "Kubot", Score: 0},
	}
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

	ps.store.RecordWin(name)
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
