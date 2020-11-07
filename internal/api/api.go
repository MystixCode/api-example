package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var db *gorm.DB

type API struct {
	Routes *Routes
}

type Routes struct {
	Root    *mux.Router
	ApiRoot *mux.Router

	Users           *mux.Router
	Groups          *mux.Router
	Scopes          *mux.Router
	ResourceServers *mux.Router
}

// Init the api1 service
func Init(root *mux.Router, database *gorm.DB) *API {
	var api API
	api.Routes = &Routes{}
	db = database

	api.Routes.Root = root
	api.Routes.ApiRoot = api.Routes.Root.PathPrefix("").Subrouter()

	api.Routes.Users = api.Routes.ApiRoot.PathPrefix("/tests").Subrouter()

	api.InitTests()

	fmt.Println("api routes loaded")
	return &api
}
