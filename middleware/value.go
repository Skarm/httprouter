package middleware

import (
	"context"
	"net/http"

	"github.com/Skarm/httprouter"
)

// WithValue is a middleware that sets a given key/value in a context chain.
func WithValue(key interface{}, val interface{}) func(next httprouter.Handle) httprouter.Handle {
	return func(next httprouter.Handle) httprouter.Handle {
		fn := func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			r = r.WithContext(context.WithValue(r.Context(), key, val))
			next(w, r, p)
		}
		return fn
	}
}
