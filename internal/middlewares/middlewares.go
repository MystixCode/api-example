package middlewares

import (
	"first_go_app/pkg/logger"

	"net/http"
)

// Json middleware to set content-type header for json
func Json(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// Logging middleware to log all requests for debugging
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Debug("\nHOST: ", r.Host, "\nMETHOD: ", r.Method, "\nURL: ", r.URL, "\nHEADERS: ", r.Header, "\nBODY: ", r.Body)
		next.ServeHTTP(w, r)
	})
}
