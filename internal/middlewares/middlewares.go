package middlewares

import (
	"fmt"
	"net/http"
)

// JSON middleware to set content-type header for json
func Json(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// Logging middleware to log all requests for debugging
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//logger.Debug(r.Method, " ", r.URL)
		fmt.Println(r.URL)
		next.ServeHTTP(w, r)
	})
}
