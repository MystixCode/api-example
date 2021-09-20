package api

import (
	"first_go_app/pkg/utils"

	"net/http"
)

func (api *API) InitTests() {
	api.Routes.Tests.HandleFunc("", getTests).Methods(http.MethodGet)
}

func getTests(w http.ResponseWriter, r *http.Request) {
	//todo: use model and unify Response via http util?
	//todo: this is just a response test
	rs := utils.Response{
		Status:  http.StatusOK,
		Message: "Hello",
		Data:    "World",
	}
	rs.Send(w)
}
