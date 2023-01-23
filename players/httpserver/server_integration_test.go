package httpserver_test

import (
	"net/http/httptest"
	"players/httpserver"
	"players/store/inmemory"
	"testing"
)

func TestIntegrationRoundtrip(t *testing.T) {
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

}
