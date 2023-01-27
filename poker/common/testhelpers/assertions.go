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

func AssertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}
func AssertNoErrorM(t *testing.T, err error, format string, args ...any) {
	t.Helper()
	if err != nil {
		t.Fatalf(format, args...)
	}
}
