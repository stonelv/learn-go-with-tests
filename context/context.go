package context

import (
	"context"
	"fmt"
	"net/http"
)

//Store is a interface defined the methods of a store
type Store interface {
	Fetch(ctx context.Context) (string, error)
}

//Server return a handler used to call fetch
func Server(store Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return
		}

		fmt.Fprint(rw, data)
	}
}
