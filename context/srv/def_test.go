package srv

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type stubSpy struct {
	response string
	t        *testing.T
	ctx      context.Context
}

func (s *stubSpy) Fetch() (string, error) {

	dataChannel := make(chan string)

	go func() {
		res := ""
		for i, letter := range s.response {
			select {
			case <-s.ctx.Done():
				return
			default:
				time.Sleep(10 * time.Millisecond)
				res += string(letter)
				fmt.Printf("fetched %d signs, now having %q\n", i, res)

			}
		}
		dataChannel <- res
	}()
	select {
	case <-s.ctx.Done():
		return "", s.ctx.Err()
	case data := <-dataChannel:
		return data, nil
	}
}

func newSpy(t *testing.T, data string, context context.Context) *stubSpy {
	s := stubSpy{ctx: context}
	s.t = t
	s.response = data
	return &s
}

func TestServer(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		data := "hello, world"

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
		spy := newSpy(t, data, request.Context())
		srv := Server(spy)

		srv.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Fatalf("got %s, want %s", response.Body.String(), data)
		}
	})
	t.Run("request with cancellation", func(t *testing.T) {
		data := "hello, data"

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(25*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)
		spy := newSpy(t, data, request.Context())
		srv := Server(spy)

		srv.ServeHTTP(response, request)

		if val := response.Body.String(); val != "" {
			t.Fatalf("expected body to be empty but received %s", val)
		}

	})

}
