package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Server configuration
type Server struct {
	Cfg    *ServerConfig
	Router *mux.Router
}

// NewServer creates a new server instance
// cfg is the server configuration
// router is the router instance
func NewServer(cfg *ServerConfig, router *mux.Router) *Server {
	return &Server{
		Cfg:    cfg,
		Router: router,
	}
}

// Starts and listens server
func (s *Server) Run() error {
	log.Println("Starting server on", s.Cfg.Addr)
	return http.ListenAndServe(s.Cfg.Addr, s.Router)
}
