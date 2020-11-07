package server

import (
	"first_go_app/internal/api"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	"fmt"
	"net/http"
)

type Server struct {
	Router   *mux.Router
	Database *gorm.DB
}

func New() *Server {
	var server Server
	server.Router = mux.NewRouter()
	return &server
}

func (s *Server) Run() {
	s.InitDatabase()
	api.Init(s.Router, s.Database)

	addr := viper.GetString("server.host") + ":" + viper.GetString("server.port")
	httpserver := &http.Server{
		Addr:    addr,
		Handler: s.Router,
	}
	fmt.Println("HTTP Server started listening on " + addr)
	fmt.Println(httpserver.ListenAndServe())
}

func (s *Server) InitDatabase() {
	var err error
	config := gorm.Config{
		Logger:      gormLogger.Default.LogMode(gormLogger.Silent),
		PrepareStmt: true,
	}
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	dbname := viper.GetString("database.name")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)
	s.Database, err = gorm.Open(mysql.Open(dsn), &config)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connected to database")
	}
}
