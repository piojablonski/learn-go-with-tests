package filesystem_test

import (
	. "players/business"
	. "players/common/testhelpers"
	. "players/store/filesystem"
	"testing"
)

func TestFilesystemStore(t *testing.T) {

	initialData := `[
{"Name": "Swiatek", "Score": 300},
{"Name": "Kubot", "Score": 0}
]`

	t.Run("get league", func(t *testing.T) {
		// src that can be read by io.Reader

		// GetAllPlayers

		database, cleanDatabase := CreateTempFile(t, initialData)
		defer cleanDatabase()
		var want = League{
			{Name: "Swiatek", Score: 300},
			{Name: "Kubot", Score: 0},
		}

		store := NewStore(database)

		got := store.GetAllPlayers()

		AssertEqual(t, got, want)

		// read again
		got = store.GetAllPlayers()

		AssertEqual(t, got, want)

	})
	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, initialData)
		defer cleanDatabase()
		want := 300
		store := NewStore(database)
		got, _ := store.GetScoreByPlayer("Swiatek")
		AssertEqual(t, got, want)
	})
	t.Run("record win for existing player", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, initialData)
		defer cleanDatabase()
		store := NewStore(database)
		store.RecordWin("Swiatek")
		got, _ := store.GetScoreByPlayer("Swiatek")
		want := 301
		AssertEqual(t, got, want)
	})
	t.Run("record win for a new player", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, initialData)
		defer cleanDatabase()
		store := NewStore(database)
		store.RecordWin("Hurkacz")
		got, _ := store.GetScoreByPlayer("Hurkacz")
		want := 1
		AssertEqual(t, got, want)
	})
}
