package middlewares

import "net/http"

// JSONHeader adds json content type header
func JSONHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(res, req)
	})
}

// HostnameHeader adds server name header
func HostnameHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("X-Server-Name", "localhost")
		next.ServeHTTP(res, req)
	})
}
