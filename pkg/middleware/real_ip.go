package middleware

import "net/http"

// headers is a string array which holds the header keys to fetch an ip address
// these values are in order of which parsing should happen.
var headers = []string{
	http.CanonicalHeaderKey("True-Client-IP"),
	http.CanonicalHeaderKey("X-Real-IP"),
	http.CanonicalHeaderKey("X-Forwarded-For"),
}

// RealIP is a middleware that sets a http.Request.RemoteAddr to the results
// of parsing either of the headers; True-Client-IP, X-Real-IP or the X-Forwarded-For.
func RealIP(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var ip string
		for _, header := range headers {
			if value := r.Header.Get(header); value != "" {
				ip = value
			}
		}
		if ip != "" {
			r.RemoteAddr = ip
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
