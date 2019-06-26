package httprouter

type middleware func(Handle) Handle

// Use appends new middleware to current Router.
func (r *Router) Use(m ...middleware) *Router {
	r.middlewares = append(m, r.middlewares...)
	return r
}
