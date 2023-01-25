package store_test

import (
	"github.com/piojablonski/learn-go-with-tests/poker/common/testhelpers"
	"github.com/piojablonski/learn-go-with-tests/poker/store"
	"io"
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
