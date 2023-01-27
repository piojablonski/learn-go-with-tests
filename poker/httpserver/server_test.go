package httpserver_test

import (
	"fmt"
	"github.com/piojablonski/learn-go-with-tests/poker/application"
	"github.com/piojablonski/learn-go-with-tests/poker/business"
	"github.com/piojablonski/learn-go-with-tests/poker/common/testhelpers"
	"github.com/piojablonski/learn-go-with-tests/poker/httpserver"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var scores = map[string]int{
	"Swiatek": 300,
	"Hurkacz": 234,
	"Kubot":   0,
}

func TestServer(t *testing.T) {
	store := &testhelpers.SpyPlayerStore{Scores: scores}
	srv := httpserver.NewPlayerServer(store)
	t.Run("return Swiatek scores", func(t *testing.T) {
		req := getPlayerScores("Swiatek")
		res := httptest.NewRecorder()
		assertStatusOk(t, res)

		srv.ServeHTTP(res, req)
		assertEqual(t, res.Body.String(), "300")
	})
	t.Run("return Hurkacz scores", func(t *testing.T) {
		req := getPlayerScores("Hurkacz")
		res := httptest.NewRecorder()
		assertStatusOk(t, res)

		srv.ServeHTTP(res, req)
		assertEqual(t, res.Body.String(), "234")

	})
	t.Run("return Kubot scores who has 0", func(t *testing.T) {
		req := getPlayerScores("Kubot")
		res := httptest.NewRecorder()
		srv.ServeHTTP(res, req)
		assertStatusOk(t, res)

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
	wantedPlayers := []business.Player{
		{Name: "Swiatek", Score: 300},
		{Name: "Hurkacz", Score: 234},
		{Name: "Kubot", Score: 0},
	}
	store := &testhelpers.SpyPlayerStore{nil, nil, wantedPlayers}
	srv := httpserver.NewPlayerServer(store)

	t.Run("it return 200 on /league", func(t *testing.T) {

		req := newLeagueRequest()
		res := httptest.NewRecorder()
		srv.ServeHTTP(res, req)

		got := getLeagueFromResponse(t, res.Result().Body)

		assertStatusOk(t, res)
		assertEqual(t, got, wantedPlayers)
	})
	t.Run("it return content type json", func(t *testing.T) {
		req := newLeagueRequest()
		res := httptest.NewRecorder()
		srv.ServeHTTP(res, req)
		// assertStatusOk(t, res.Result().StatusCode)
		assertContentType(t, res)
	})

	t.Run("it returns 200 when hit /game", func(t *testing.T) {

		req, _ := newGameRequest()
		res := httptest.NewRecorder()

		srv.ServeHTTP(res, req)
		assertStatusOk(t, res)
	})
}

func newGameRequest() (*http.Request, error) {
	return http.NewRequest(http.MethodGet, "/game", nil)
}

func assertContentType(t *testing.T, res *httptest.ResponseRecorder) {
	t.Helper()
	want := httpserver.ApplicationJsonContentType
	got := res.Header().Get("content-type")

	assertEqual(t, got, want)
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func getLeagueFromResponse(t *testing.T, body io.Reader) []business.Player {
	t.Helper()
	got, err := application.ReadPlayers(body)
	if err != nil {
		t.Fatal(err)
	}
	return got
}

func TestStoreWins(t *testing.T) {
	store := &testhelpers.SpyPlayerStore{Scores: scores}
	srv := httpserver.NewPlayerServer(store)

	req := postPlayerScores("Radwańska")
	res := httptest.NewRecorder()
	srv.ServeHTTP(res, req)

	assertStatusOk(t, res)

	if len(store.WinCalls) != 1 {
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

func assertStatusOk(t *testing.T, response *httptest.ResponseRecorder) {
	t.Helper()
	if response.Code != http.StatusOK {
		t.Fatalf("expected to receive status 200 but received %d", response.Code)
	}
}
func assertEqual(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got: %v, want: %v", got, want)
	}
}
