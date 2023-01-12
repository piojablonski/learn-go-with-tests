package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people with names", func(t *testing.T) {
		got := Hello("Piotr")
		want := "Hello Piotr!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to people without a name", func(t *testing.T) {
		got := Hello("")
		want := "Hello World!"
		assertCorrectMessage(t, got, want)

	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
