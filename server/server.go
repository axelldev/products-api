package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type routeBinder func(Server, *mux.Router)

type Config struct {
	Port        string
	JWTSecret   string
	DatabaseUrl string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	conifg *Config
	router *mux.Router
}

func (b *Broker) Config() *Config {
	return b.conifg
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("Port is required.")
	}

	if config.JWTSecret == "" {
		return nil, errors.New("JWT secret is required.")
	}

	if config.DatabaseUrl == "" {
		return nil, errors.New("Database url is required.")
	}

	return &Broker{
		conifg: config,
		router: mux.NewRouter(),
	}, nil
}

func (b *Broker) Run(binder routeBinder) {
	binder(b, b.router)
	log.Printf("Server started on port %s", b.conifg.Port)
	if err := http.ListenAndServe(b.conifg.Port, b.router); err != nil {
		log.Fatal("Error starting the server:", err)
	}
}
