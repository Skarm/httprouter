package middleware

import (
	"net/http"
	"strings"

	"github.com/Skarm/httprouter"
)

// Heartbeat endpoint middleware useful to setting up a path like
// `/ping` that load balancers or uptime testing external services
// can make a request before hitting any routes. It's also convenient
// to place this above ACL middlewares as well.
func Heartbeat(endpoint string) func(httprouter.Handle) httprouter.Handle {
	f := func(h httprouter.Handle) httprouter.Handle {
		fn := func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			if r.Method == http.MethodGet && strings.EqualFold(r.URL.Path, endpoint) {
				w.Header().Set("Content-Type", "text/plain")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("."))
				return
			}
			h(w, r, p)
		}
		return fn
	}
	return f
}
