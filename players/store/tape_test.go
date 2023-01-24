package store_test

import (
	"io"
	"players/common/testhelpers"
	"players/store"
	"testing"
)

func TestTape_Write(t *testing.T) {
	file, clean := testhelpers.CreateTempFile(t, "1234567")
	defer clean()

	tape := store.Tape{File: file}

	tape.Write([]byte("abc"))

	file.Seek(0, io.SeekStart)
	fcontent, _ := io.ReadAll(file)
	got := string(fcontent)
	want := "abc"
	testhelpers.AssertEqual(t, got, want)
}
