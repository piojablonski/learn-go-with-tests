package testhelpers

import (
	"net/http"
	"reflect"
	"testing"
)

func AssertStatusOk(t *testing.T, code int) {
	t.Helper()
	if code != http.StatusOK {
		t.Fatalf("expected to receive status 200 but received %d", code)
	}
}
func AssertEqual(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got: %v, want: %v", got, want)
	}
}
