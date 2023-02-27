package middleware

import "github.com/go-chi/chi/v5/middleware"

// NewWrapResponseWriter wraps an http.ResponseWriter returning a proxy that allows you to
// hook into various parts of the response process.
//
// This is a wrapper around go-chi/middleware/NewWrapResponseWriter.
var NewWrapResponseWriter = middleware.NewWrapResponseWriter
