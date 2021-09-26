package api

import (
	"first_go_app/pkg/errors"
	"first_go_app/pkg/logger"
	"first_go_app/pkg/models"
	"first_go_app/pkg/utils"
	"fmt"

	"gorm.io/gorm"

	"net/http"
)

func (api *API) InitUsers() {
	api.Routes.Users.HandleFunc("", getUsers).Methods(http.MethodGet)
	//	api.Routes.Users.HandleFunc("", createUser).Methods(http.MethodPost)
	//	api.Routes.Users.HandleFunc("/{id}", getUser).Methods(http.MethodGet)
	//	api.Routes.Users.HandleFunc("/{id}", updateUser).Methods(http.MethodPut)
	//	api.Routes.Users.HandleFunc("/{id}", deleteUser).Methods(http.MethodDelete)

	//	api.Routes.Users.HandleFunc("/{id}/groups", getUserGroups).Methods(http.MethodGet)
}

func getUsers(w http.ResponseWriter, r *http.Request) {

	var user models.User

	users, err := user.GetAll(db)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.SendErrorResponse(w, http.StatusBadRequest, err.Error())
		} else {
			fmt.Println(err)
			logger.Debug(err)
			utils.SendErrorResponse(w, http.StatusInternalServerError, errors.ServerError.Error())
		}
		return
	}

	if len(*users) == 0 {
		utils.SendErrorResponse(w, http.StatusNotFound, "no users found")
		return
	}

	rs := utils.Response{
		Status:  http.StatusOK,
		Message: "users found",
		Data:    users,
	}
	rs.Send(w)
}
