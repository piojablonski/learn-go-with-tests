package srv

import (
	"net/http"
)

type Store interface {
	Fetch() (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		val, _ := store.Fetch()
		w.Write([]byte(val))
	}
}
