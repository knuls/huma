package middleware

import "github.com/go-chi/chi/v5/middleware"

// NoCache is a middleware that prevents clients from caching.
//
// This is a wrapper around go-chi/middleware/NoCache.
var NoCache = middleware.NoCache
