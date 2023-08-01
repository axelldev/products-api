package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	// Server configuration
	Cfg    *ServerConfig
	Router *mux.Router
}

func NewServer(cfg *ServerConfig, router *mux.Router) *Server {
	// Create a new server instance
	return &Server{
		Cfg:    cfg,
		Router: router,
	}
}

func (s *Server) Run() error {
	// Start the server
	log.Println("Starting server on", s.Cfg.Addr)
	return http.ListenAndServe(s.Cfg.Addr, s.Router)
}
