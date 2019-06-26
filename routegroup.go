package httprouter

import (
	"net/http"
)

// Router is a http.Handler that wraps httprouter.Router with additional features.
type RouteGroup struct {
	middlewares []middleware
	path        string
	router      *Router
}

// New returns *RouteGroup with a new initialized *Router embedded.
func newRouteGroup(r *Router, path string, m ...middleware) *RouteGroup {
	if path[0] != '/' {
		panic("path must begin with '/' in path '" + path + "'")
	}

	//Strip traling / (if present) as all added sub paths must start with a /
	if path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}

	return &RouteGroup{
		middlewares: append(m, r.middlewares...),
		path:        path,
		router:      r,
	}
}

func (r *RouteGroup) Group(path string, m ...middleware) *RouteGroup {
	return newRouteGroup(r.router, r.subPath(path), m...)
}

// Handle registers a new request handle combined with middlewares.
func (r *RouteGroup) Handle(method, path string, handle Handle) {
	r.router.Handle(method, r.subPath(path), handle)
}

// Handler is an adapter for http.Handler.
func (r *RouteGroup) Handler(method, path string, handler http.Handler) {
	r.router.Handler(method, r.subPath(path), handler)
}

// HandlerFunc is an adapter for http.HandlerFunc.
func (r *RouteGroup) HandlerFunc(method, path string, handler http.HandlerFunc) {
	r.router.HandlerFunc(method, r.subPath(path), handler)
}

// Use appends new middleware to current RouteGroup.
func (r *RouteGroup) Use(m ...middleware) *RouteGroup {
	r.middlewares = append(m, r.middlewares...)
	return r
}

// GET is a shortcut for Router.Handle("GET", path, handle)
func (r *RouteGroup) GET(path string, handle Handle) {
	r.Handle("GET", path, handle)
}

// HEAD is a shortcut for Router.Handle("HEAD", path, handle)
func (r *RouteGroup) HEAD(path string, handle Handle) {
	r.Handle("HEAD", path, handle)
}

// OPTIONS is a shortcut for Router.Handle("OPTIONS", path, handle)
func (r *RouteGroup) OPTIONS(path string, handle Handle) {
	r.Handle("OPTIONS", path, handle)
}

// POST is a shortcut for Router.Handle("POST", path, handle)
func (r *RouteGroup) POST(path string, handle Handle) {
	r.Handle("POST", path, handle)
}

// PUT is a shortcut for Router.Handle("PUT", path, handle)
func (r *RouteGroup) PUT(path string, handle Handle) {
	r.Handle("PUT", path, handle)
}

// PATCH is a shortcut for Router.Handle("PATCH", path, handle)
func (r *RouteGroup) PATCH(path string, handle Handle) {
	r.Handle("PATCH", path, handle)
}

// DELETE is a shortcut for Router.Handle("DELETE", path, handle)
func (r *RouteGroup) DELETE(path string, handle Handle) {
	r.Handle("DELETE", path, handle)
}

func (r *RouteGroup) subPath(path string) string {
	if path[0] != '/' {
		panic("path should start with '/' in path '" + path + "'.")
	}
	return r.path + path
}
