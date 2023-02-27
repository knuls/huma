package middleware

import "github.com/go-chi/chi/v5/middleware"

// RequestID is a middleware that sets a request ID into the request context.
//
// This is a wrapper around go-chi/middleware/RequestID.
var RequestID = middleware.RequestID

// GetReqID returns a request ID from the given context if one is present.
// Returns the empty string if a request ID cannot be found.
//
// This is a wrapper around go-chi/middleware/GetReqID.
var GetReqID = middleware.GetReqID
