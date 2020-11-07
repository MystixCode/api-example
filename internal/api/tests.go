package api

import (
	"encoding/json"
	"net/http"
)

func (api *API) InitTests() {
	api.Routes.Users.HandleFunc("", getTests).Methods(http.MethodGet)
	//todo: more routes here
}

func getTests(w http.ResponseWriter, r *http.Request) {
	//todo: use model and unify Response via http util?
	//todo: this is just a response test
	bytes, _ := json.Marshal(r)
	w.Write(bytes)
}
