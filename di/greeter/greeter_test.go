package greeter

import (
	"bytes"
	"testing"
)

func TestGreeter(t *testing.T) {
	var buffer bytes.Buffer
	Greet(&buffer, "Piotr")
	got := buffer.String()
	want := "Hello Piotr"

	if got != want {
		t.Errorf("got: %s, want %s", got, want)
	}

}
