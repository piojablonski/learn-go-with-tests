package testhelpers

import (
	"os"
	"testing"
)

func CreateTempFile(t testing.TB, initialData string) (*os.File, func()) {
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
