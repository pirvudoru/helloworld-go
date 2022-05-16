package http_server

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(c context.Context) (string, error)
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return
		}

		_, _ = fmt.Fprint(w, data)
	}
}
