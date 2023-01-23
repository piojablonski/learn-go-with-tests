package countdown

import (
	"bytes"
	"reflect"
	"testing"
)

type spySleeper struct {
	calls  []string
	buffer *bytes.Buffer
}

func (s *spySleeper) Sleep() {
	s.calls = append(s.calls, "sleep")
}

func (s *spySleeper) Write(b []byte) (int, error) {
	s.calls = append(s.calls, "write")
	return s.buffer.Write(b)
}

func TestCountdown(t *testing.T) {

	t.Run("render", func(t *testing.T) {
		testHelper := &spySleeper{buffer: &bytes.Buffer{}}
		Countdown(testHelper, testHelper)
		want := `3
2
1
go!`
		got := testHelper.buffer.String()

		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})
	t.Run("order", func(t *testing.T) {
		testHelper := &spySleeper{buffer: &bytes.Buffer{}}
		Countdown(testHelper, testHelper)

		want := []string{"write", "sleep", "write", "sleep", "write", "sleep", "write"}
		got := testHelper.calls

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want: %v", testHelper.calls, want)
		}
	})

}
