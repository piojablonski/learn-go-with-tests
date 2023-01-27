package httpserver

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/piojablonski/learn-go-with-tests/poker/business"
	"github.com/piojablonski/learn-go-with-tests/poker/store"
	"html/template"
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
	router.HandleFunc("/game", gameHandler)
	srv.store = store
	return srv

}

const ApplicationJsonContentType = "application/json"

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

func gameHandler(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFS(postTemplates, "templates/*")
	if err != nil {
		http.Error(w, fmt.Sprintf("problem loading template %s", err.Error()), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("problem executing template %s", err.Error()), http.StatusInternalServerError)
		return

	}

}

func (ps *PlayerServer) leagueHandler(w http.ResponseWriter, _ *http.Request) {
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
