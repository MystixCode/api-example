package api

import (
	"first_go_app/pkg/errors"
	"first_go_app/pkg/models"
	"first_go_app/pkg/utils"

	"encoding/json"
	"net/http"
)

func (api *API) InitTests() {
	api.Routes.Tests.HandleFunc("", getTests).Methods(http.MethodGet)
	api.Routes.Tests.HandleFunc("", createTest).Methods(http.MethodPost)
	api.Routes.Tests.HandleFunc("/{id}", getTest).Methods(http.MethodGet)
	api.Routes.Tests.HandleFunc("/{id}", updateTest).Methods(http.MethodPut)
	api.Routes.Tests.HandleFunc("/{id}", deleteTest).Methods(http.MethodDelete)
}

func getTests(w http.ResponseWriter, r *http.Request) {
	var test models.Test

	tests, err := test.GetAll(db)
	if err != nil {
		if err == errors.RecordNotFound {
			utils.SendErrorResponse(w, http.StatusBadRequest, err.Error())
		} else {
			utils.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if len(*tests) == 0 {
		utils.SendErrorResponse(w, http.StatusNotFound, "no tests found")
		return
	}

	rs := utils.Response{
		Status:  http.StatusOK,
		Message: "tests found",
		Data:    tests,
	}
	rs.Send(w)
}

func createTest(w http.ResponseWriter, r *http.Request) {
	var test models.Test
	var newTest models.TestBeforeSave

	err := json.NewDecoder(r.Body).Decode(&newTest)
	if newTest == (models.TestBeforeSave{}) {
		utils.SendErrorResponse(w, http.StatusBadRequest, errors.InvalidData.Error())
		return
	}
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, errors.InvalidData.Error())
		return
	}

	err = test.Save(db, &newTest)
	if err != nil {
		if err == errors.InvalidData {
			utils.SendErrorResponse(w, http.StatusBadRequest, err.Error())
		} else {
			utils.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	rs := utils.Response{
		Status:  http.StatusCreated,
		Message: "test created successfully",
		Data:    test,
	}
	rs.Send(w)
}

func getTest(w http.ResponseWriter, r *http.Request) {
	var test models.Test

	id := utils.GetMuxParam(r, "id")

	if id == "" {
		utils.SendErrorResponse(w, http.StatusBadRequest, "invalid id")
		return
	}

	err := test.GetByID(db, id)
	if err != nil {
		if err == errors.RecordNotFound {
			utils.SendErrorResponse(w, http.StatusBadRequest, err.Error())
		} else {
			utils.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	response := utils.Response{
		Status:  http.StatusOK,
		Message: "test found",
		Data:    test,
	}
	response.Send(w)
}

func updateTest(w http.ResponseWriter, r *http.Request) {
	var test models.Test
	var newTest models.TestBeforeSave

	id := utils.GetMuxParam(r, "id")
	if id == "" {
		utils.SendErrorResponse(w, http.StatusBadRequest, "invalid id")
		return
	}

	err := json.NewDecoder(r.Body).Decode(&newTest)
	if newTest == (models.TestBeforeSave{}) {
		utils.SendErrorResponse(w, http.StatusBadRequest, errors.InvalidData.Error())
		return
	}
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, errors.InvalidData.Error())
		return
	}

	err = test.Update(db, id, &newTest)
	if err != nil {
		if err == errors.RecordNotFound {
			utils.SendErrorResponse(w, http.StatusNotFound, err.Error())
		}
		if err == errors.InvalidData {
			utils.SendErrorResponse(w, http.StatusBadRequest, err.Error())
		} else {
			utils.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	response := utils.Response{
		Status:  http.StatusOK,
		Message: "test updated successfully",
		Data:    test,
	}
	response.Send(w)
}

func deleteTest(w http.ResponseWriter, r *http.Request) {
	var test models.Test

	id := utils.GetMuxParam(r, "id")
	if id == "" {
		utils.SendErrorResponse(w, http.StatusBadRequest, "invalid id")
		return
	}

	err := test.GetByID(db, id)
	if err != nil {
		if err == errors.RecordNotFound {
			utils.SendErrorResponse(w, http.StatusBadRequest, err.Error())
		} else {
			utils.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	err = test.Delete(db, test.ID)
	if err != nil {
		if err == errors.RecordNotFound {
			utils.SendErrorResponse(w, http.StatusBadRequest, err.Error())
		} else {
			utils.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	response := utils.Response{
		Status:  http.StatusOK,
		Message: "test deleted",
		Data:    test,
	}
	response.Send(w)
}
