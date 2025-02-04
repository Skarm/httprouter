package middleware

import (
	"net/http"
	"strings"

	"github.com/Skarm/httprouter"
)

// SetHeader is a convenience handler to set a response header key/value
func SetHeader(key, value string) func(next httprouter.Handle) httprouter.Handle {
	return func(next httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			w.Header().Set(key, value)
			next(w, r, p)
		}
	}
}

// AllowContentType enforces a whitelist of request Content-Types otherwise responds
// with a 415 Unsupported Media Type status.
func AllowContentType(contentTypes ...string) func(next httprouter.Handle) httprouter.Handle {
	cT := []string{}
	for _, t := range contentTypes {
		cT = append(cT, strings.ToLower(t))
	}

	return func(next httprouter.Handle) httprouter.Handle {
		fn := func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			if r.ContentLength == 0 {
				// skip check for empty content body
				next(w, r, p)
				return
			}

			s := strings.ToLower(strings.TrimSpace(r.Header.Get("Content-Type")))
			if i := strings.Index(s, ";"); i > -1 {
				s = s[0:i]
			}

			for _, t := range cT {
				if t == s {
					next(w, r, p)
					return
				}
			}

			w.WriteHeader(http.StatusUnsupportedMediaType)
		}
		return fn
	}
}
