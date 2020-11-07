package server

import (
	"first_go_app/internal/api"
)

func (s *Server) initServices() {
	api.Init(s.Router, s.Database)
}
