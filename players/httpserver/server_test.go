package httpserver_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	. "players/business"
	"players/httpserver"
	"reflect"
	"testing"
)

var scores = map[string]int{
	"Swiatek": 300,
	"Hurkacz": 234,
	"Kubot":   0,
}

type stubPlayerStore struct {
	scores     map[string]int
	operations []string
	league     []Player
}

func (s *stubPlayerStore) GetScoreByPlayer(name string) (score int, found bool) {
	score, found = s.scores[name]
	return
}

func (s *stubPlayerStore) RecordWin(name string) {
	s.operations = append(s.operations, "record")
}

func TestServer(t *testing.T) {
	store := &stubPlayerStore{scores: scores}
	srv := httpserver.NewPlayerServer(store)
	t.Run("return Swiatek scores", func(t *testing.T) {
		req := getPlayerScores("Swiatek")
		res := httptest.NewRecorder()
		assertStatusOk(t, res.Code)

		srv.ServeHTTP(res, req)
		assertEqual(t, res.Body.String(), "300")
	})
	t.Run("return Hurkacz scores", func(t *testing.T) {
		req := getPlayerScores("Hurkacz")
		res := httptest.NewRecorder()
		assertStatusOk(t, res.Code)

		srv.ServeHTTP(res, req)
		assertEqual(t, res.Body.String(), "234")

	})
	t.Run("return Kubot scores who has 0", func(t *testing.T) {
		req := getPlayerScores("Kubot")
		res := httptest.NewRecorder()
		srv.ServeHTTP(res, req)
		assertStatusOk(t, res.Code)

		assertEqual(t, res.Body.String(), "0")
		if res.Code != http.StatusOK {
			t.Fatalf("expected to receive status 200 but received %d", res.Code)
		}

	})
	t.Run("record an operation", func(t *testing.T) {
		req := getPlayerScores("Fręch")
		res := httptest.NewRecorder()
		srv.ServeHTTP(res, req)
		if res.Code != http.StatusNotFound {
			t.Fatalf("expected to receive status 404 but received %d", res.Code)
		}

	})
}

func TestLeague(t *testing.T) {
	wantedPlayers := []Player{
		{Name: "Swiatek", Score: 300},
		{Name: "Hurkacz", Score: 234},
		{Name: "Kubot", Score: 0},
	}
	store := &stubPlayerStore{nil, nil, wantedPlayers}
	srv := httpserver.NewPlayerServer(store)

	t.Run("it return 200 on /league", func(t *testing.T) {

		req, _ := http.NewRequest(http.MethodGet, "/league", nil)
		res := httptest.NewRecorder()
		srv.ServeHTTP(res, req)

		var got []Player

		err := json.NewDecoder(res.Body).Decode(&got)
		if err != nil {
			t.Fatalf("unable to decode got %q", err)
		}

		if !reflect.DeepEqual(got, wantedPlayers) {
			t.Errorf("wanted %v got %v", wantedPlayers, got)

		}
		assertStatusOk(t, res.Code)
	})
}

func TestStoreWins(t *testing.T) {
	store := &stubPlayerStore{scores: scores}
	srv := httpserver.NewPlayerServer(store)

	req := postPlayerScores("Radwańska")
	res := httptest.NewRecorder()
	srv.ServeHTTP(res, req)

	assertStatusOk(t, res.Code)

	if store.operations[0] != "record" {
		t.Fatalf("expected record operation")
	}
}

func getPlayerScores(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/player/%s", name), nil)
	return req
}
func postPlayerScores(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/player/%s", name), nil)
	return req
}

func assertStatusOk(t *testing.T, code int) {
	t.Helper()
	if code != http.StatusOK {
		t.Fatalf("expected to receive status 200 but received %d", code)
	}
}
func assertEqual(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}
