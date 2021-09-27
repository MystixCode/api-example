package api

import (
	"encoding/json"
	"net/http"
)

func (api *API) InitIndex() {
	api.Routes.Index.HandleFunc("/", getIndex).Methods(http.MethodGet)
}

func getIndex(w http.ResponseWriter, r *http.Request) {
	msg := struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}{
		Status:  http.StatusOK,
		Message: "api-example root / endpoint",
	}

	js, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
