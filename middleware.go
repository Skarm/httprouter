package httprouter

import (
	"context"
	"net/http"
)

type middleware func(Handle) Handle

// Use appends new middleware to current Router.
func (r *Router) Use(m ...middleware) *Router {
	r.middlewares = append(m, r.middlewares...)
	return r
}

// Wrap a given handle with the current stack from the result of Use() calls.
func (r *Router) Wrap(fn Handle) Handle {
	l := len(r.middlewares)
	if l == 0 {
		return fn
	}

	// There is at least one item in the list. Starting
	// with the last item, create the handler to be
	// returned:
	var result Handle
	result = r.middlewares[l-1](fn)

	// Reverse through the stack for the remaining elements,
	// and wrap the result with each layer:
	for i := 0; i < (l - 1); i++ {
		result = r.middlewares[l-(2+i)](result)
	}

	return result
}

// WrapHandler wraps http.Handler and returns httprouter.Handle
func WrapHandler(next http.Handler) Handle {
	return func(w http.ResponseWriter, r *http.Request, p Params) {
		//pass httprouter.Params to request context
		ctx := context.WithValue(r.Context(), "params", p)
		//call next middleware with new context
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// FetchParams from http.Request
func FetchParams(req *http.Request) Params {
	ctx := req.Context()
	return ctx.Value("params").(Params)
}
