package main

import (
	"context"
	"log"
	"os"

	"github.com/axelldev/products-api/handlers"
	"github.com/axelldev/products-api/server"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	port := ":" + os.Getenv("PORT")
	JWTSecret := os.Getenv("JWT_SECRET")
	databaseURL := os.Getenv("DATABASE_URL")

	svr, err := server.NewServer(context.Background(), &server.Config{
		Port:        port,
		JWTSecret:   JWTSecret,
		DatabaseUrl: databaseURL,
	})

	if err != nil {
		log.Fatal(err)
	}

	svr.Run(RouteBinder)
}

func RouteBinder(s server.Server, r *mux.Router) {
	r.HandleFunc("/", handlers.HomeHandler(s))
}
