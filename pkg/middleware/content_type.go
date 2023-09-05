package middleware

import "github.com/go-chi/chi/v5/middleware"

// AllowContentType is a middleware that white lists request accepted content types.
//
// This is a wrapper around go-chi/middleware/AllowContentType.
var AllowContentType = middleware.AllowContentType
