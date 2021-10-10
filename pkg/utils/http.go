package utils

import (
	"fmt"

	"github.com/gorilla/mux"

	"encoding/json"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r *Response) Send(w http.ResponseWriter) {
	w.WriteHeader(r.Status)
	bytes, _ := json.Marshal(r)
	_, err := w.Write(bytes)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, nil)
	}
}

func SendErrorResponse(w http.ResponseWriter, status int, error interface{}) {
	w.WriteHeader(status)
	var response struct {
		Status int         `json:"status"`
		Error  interface{} `json:"error"`
	}
	response.Status = status
	response.Error = error
	bytes, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}

	_, err = w.Write(bytes)
	if err != nil {
		fmt.Println(err)
	}
}

func GetMuxParam(r *http.Request, index string) string {
	params := mux.Vars(r)
	return params[index]
}
