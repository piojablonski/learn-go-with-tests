package store_test

import (
	"io"
	"os"
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

		database, cleanDatabase := createTempFile(t, initialData)
		defer cleanDatabase()
		var want = League{
			{Name: "Swiatek", Score: 300},
			{Name: "Kubot", Score: 0},
		}

		store := FilesystemStore{database}

		got := store.GetAllPlayers()

		AssertEqual(t, got, want)

		// read again
		got = store.GetAllPlayers()

		AssertEqual(t, got, want)

	})
	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, initialData)
		defer cleanDatabase()
		want := 300
		store := FilesystemStore{database}
		got, _ := store.GetScoreByPlayer("Swiatek")
		AssertEqual(t, got, want)
	})
	t.Run("record win", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, initialData)
		defer cleanDatabase()
		store := FilesystemStore{database}
		store.RecordWin("Swiatek")
		got, _ := store.GetScoreByPlayer("Swiatek")
		want := 301
		AssertEqual(t, got, want)
	})
}

func createTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()
	f, err := os.CreateTemp("", "db1")
	if err != nil {
		t.Fatalf("problem creating temporary file %q", err)
	}

	f.WriteString(initialData)

	removeFile := func() {
		err = os.Remove(f.Name())
		if err != nil {
			t.Fatalf("problem removing temporary file %q", err)
		}
	}

	return f, removeFile
}
