package middleware

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/Skarm/httprouter"
	"golang.org/x/net/context"
)

// Key to use when setting the request ID.
type ctxKeyRequestID int

const (
	// RequestIDKey is the key that holds the unique request ID in a request context.
	RequestIDKey ctxKeyRequestID = 0
	// RequestIDHeaderKey header name
	RequestIDHeaderKey = "X-Request-ID"
)

// RequestID Middleware.
// Checks the X-Request-ID header. If not found,
// generates a new uuid, and inserts whichever
// on the context before calling the next function.
// Should generally be the outermost middleware, so that
// all other middlewares have a request id available.
func RequestID(fn httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		reqID := r.Header.Get(RequestIDHeaderKey)
		if reqID == "" {
			reqID = uuid.New().String()
		}

		ctx := context.WithValue(r.Context(), RequestIDKey, reqID)
		fn(w, r.WithContext(ctx), p)
	}
}

// GetReqID returns a request ID from the given context if one is present.
// Returns the empty string if a request ID cannot be found.
func GetReqID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if reqID, ok := ctx.Value(RequestIDKey).(string); ok {
		return reqID
	}
	return ""
}