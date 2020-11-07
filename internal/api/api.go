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

	Index *mux.Router
	Tests *mux.Router
}

// Init the api1 service
func Init(root *mux.Router, database *gorm.DB) *API {
	var api API
	api.Routes = &Routes{}
	db = database

	api.Routes.Root = root
	api.Routes.ApiRoot = api.Routes.Root.PathPrefix("").Subrouter()
	api.Routes.Index = api.Routes.ApiRoot.PathPrefix("/").Subrouter()
	api.Routes.Tests = api.Routes.ApiRoot.PathPrefix("/tests").Subrouter()

	api.InitIndex()
	api.InitTests()

	fmt.Println("api routes loaded")
	return &api
}
