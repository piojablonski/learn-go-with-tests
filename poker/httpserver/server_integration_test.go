package httpserver_test

import (
	"github.com/piojablonski/learn-go-with-tests/poker/business"
	"github.com/piojablonski/learn-go-with-tests/poker/common/testhelpers"
	"github.com/piojablonski/learn-go-with-tests/poker/httpserver"
	"github.com/piojablonski/learn-go-with-tests/poker/store/filesystem"
	"net/http/httptest"
	"testing"
)

func TestIntegrationRoundtrip(t *testing.T) {
	t.Run("get score", func(t *testing.T) {
		db, clean := testhelpers.CreateTempFile(t, "[]")
		defer clean()
		store, err := filesystem.NewStore(db)
		testhelpers.AssertNoError(t, err)
		srv := httpserver.NewPlayerServer(store)
		srv.ServeHTTP(httptest.NewRecorder(), postPlayerScores("Swiatek"))
		srv.ServeHTTP(httptest.NewRecorder(), postPlayerScores("Swiatek"))
		srv.ServeHTTP(httptest.NewRecorder(), postPlayerScores("Swiatek"))

		res := httptest.NewRecorder()
		srv.ServeHTTP(res, getPlayerScores("Swiatek"))

		want := "3"
		got := res.Body.String()

		assertStatusOk(t, res)

		assertEqual(t, got, want)
	})

	t.Run("get League", func(t *testing.T) {
		db, clean := testhelpers.CreateTempFile(t, "[]")
		defer clean()
		var store, err = filesystem.NewStore(db)
		testhelpers.AssertNoError(t, err)
		srv := httpserver.NewPlayerServer(store)
		srv.ServeHTTP(httptest.NewRecorder(), postPlayerScores("Swiatek"))
		srv.ServeHTTP(httptest.NewRecorder(), postPlayerScores("Swiatek"))

		res := httptest.NewRecorder()
		srv.ServeHTTP(res, newLeagueRequest())
		assertStatusOk(t, res)

		got := getLeagueFromResponse(t, res.Body)
		want := []business.Player{
			{Name: "Swiatek", Score: 2},
		}

		assertEqual(t, got, want)

	})

	t.Run("works with an empty file", func(t *testing.T) {
		db, clean := testhelpers.CreateTempFile(t, "")
		defer clean()
		var _, err = filesystem.NewStore(db)
		testhelpers.AssertNoError(t, err)
	})

}
