package server

import (
	"first_go_app/internal/api"
	"first_go_app/internal/middlewares"
	"first_go_app/pkg/logger"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	"fmt"
	"net/http"
)

type Server struct {
	HttpServer *http.Server
	Router     *mux.Router
	Database   *gorm.DB
	//	Api			*api.API
}

func New() *Server {
	var s Server
	s.Router = mux.NewRouter()
	return &s
}

func (s *Server) Run() {
	s.InitDatabase()
	s.InitApi()
	s.InitRouter()
	s.InitHttpServer()

	logger.Info("HTTP Server started listening on ", s.HttpServer.Addr)
	logger.Fatal(s.HttpServer.ListenAndServe())
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
		logger.Fatal(err)
	} else {
		logger.Info("Database initialised")
	}
}

func (s *Server) InitApi() {
	api.Init(s.Router, s.Database)
	logger.Info("Api initialised")
}

func (s *Server) InitRouter() {
	s.Router.Use(middlewares.Json)
	s.Router.NotFoundHandler = http.HandlerFunc(Custom404)
	s.Router.MethodNotAllowedHandler = http.HandlerFunc(Custom405)
	//if logger.debug == true use middlewares.Logging
	if viper.GetBool("logger.debug") {
		s.Router.Use(middlewares.Logging)
		logger.Debug("Debug Mode enabled")
	}
	logger.Info("router initialised")
}

func (s *Server) InitHttpServer() {
	addr := viper.GetString("server.host") + ":" + viper.GetString("server.port")
	s.HttpServer = &http.Server{
		Addr:    addr,
		Handler: s.Router,
	}
	logger.Info("HttpServer initialised")
}
