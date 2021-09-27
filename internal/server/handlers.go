package server

import (
	"api-example/pkg/utils"

	"net/http"
)

// Custom404 return 404 in json format
func Custom404(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	utils.SendErrorResponse(w, http.StatusNotFound, "no endpoint found")
}

// Custom405 return 405 in json format
func Custom405(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	utils.SendErrorResponse(w, http.StatusMethodNotAllowed, "the used http method is not allowed on this endpoint")
}
