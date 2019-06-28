package middleware

import (
	"expvar"
	"fmt"
	"net/http"
	"net/http/pprof"

	"github.com/Skarm/httprouter"
)

// Profiler is a convenient subrouter used for mounting net/http/pprof. ie.
//
//  func MyService() http.Handler {
//	  r := httprouter.New()
//    // ..middlewares
//    r.Mount("/debug", middleware.Profiler())
//    // ..routes
//    return r
//  }
func Profiler() http.Handler {
	r := httprouter.New()
	r.Use(NoCache)

	r.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		http.Redirect(w, r, r.RequestURI+"/pprof/", 301)
	})

	r.Handle(http.MethodGet, "/pprof", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		http.Redirect(w, r, r.RequestURI+"/", 301)
	})

	r.HandlerFunc(http.MethodGet, "/pprof/*", pprof.Index)
	r.HandlerFunc(http.MethodGet, "/pprof/cmdline", pprof.Cmdline)
	r.HandlerFunc(http.MethodGet, "/pprof/profile", pprof.Profile)
	r.HandlerFunc(http.MethodGet, "/pprof/symbol", pprof.Symbol)
	r.HandlerFunc(http.MethodGet, "/pprof/trace", pprof.Trace)
	r.Handle(http.MethodGet, "/vars", expVars)

	return r
}

// Replicated from expvar.go as not public.
func expVars(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	first := true
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{\n")
	expvar.Do(func(kv expvar.KeyValue) {
		if !first {
			fmt.Fprintf(w, ",\n")
		}
		first = false
		fmt.Fprintf(w, "%q: %s", kv.Key, kv.Value)
	})
	fmt.Fprintf(w, "\n}\n")
}
