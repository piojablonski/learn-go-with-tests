package httpserver_test

import (
	"net/http/httptest"
	"players/business"
	"players/httpserver"
	"players/store/inmemory"
	"testing"
)

func TestIntegrationRoundtrip(t *testing.T) {
	t.Run("get score", func(t *testing.T) {
		var store = inmemory.NewInmemoryPlayerStore()
		srv := httpserver.NewPlayerServer(store)
		srv.ServeHTTP(httptest.NewRecorder(), postPlayerScores("Swiatek"))
		srv.ServeHTTP(httptest.NewRecorder(), postPlayerScores("Swiatek"))
		srv.ServeHTTP(httptest.NewRecorder(), postPlayerScores("Swiatek"))

		res := httptest.NewRecorder()
		srv.ServeHTTP(res, getPlayerScores("Swiatek"))

		want := "3"
		got := res.Body.String()

		assertStatusOk(t, res.Code)

		assertEqual(t, got, want)
	})

	t.Run("get League", func(t *testing.T) {
		var store = inmemory.NewInmemoryPlayerStore()
		srv := httpserver.NewPlayerServer(store)
		srv.ServeHTTP(httptest.NewRecorder(), postPlayerScores("Swiatek"))
		srv.ServeHTTP(httptest.NewRecorder(), postPlayerScores("Swiatek"))

		res := httptest.NewRecorder()
		srv.ServeHTTP(res, newLeagueRequest())
		assertStatusOk(t, res.Code)

		got := getLeagueFromResponse(t, res.Body)
		want := []business.Player{
			{Name: "Swiatek", Score: 2},
		}

		assertEqual(t, got, want)

	})

}
