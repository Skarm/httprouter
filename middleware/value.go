package middleware

import (
	"context"
	"net/http"

	"github.com/Skarm/httprouter"
)

// WithValue is a middleware that sets a given key/value in a context chain.
func WithValue(key interface{}, val interface{}) func(next httprouter.Handler) httprouter.Handler {
	return func(next httprouter.Handler) httprouter.Handler {
		fn := func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			r = r.WithContext(context.WithValue(r.Context(), key, val))
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
