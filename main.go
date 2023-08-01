package main

import (
	"net/http"

	"github.com/axelldev/products-api/handler"
	"github.com/axelldev/products-api/response"
	"github.com/axelldev/products-api/server"
	"github.com/gorilla/mux"
)

func main() {
	cfg := server.NewServerConfig()

	r := mux.NewRouter()
	r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		response.Json(w, http.StatusOK, map[string]bool{"ok": true})
	})
	r.HandleFunc("/api/products", handler.HandleProducts)

	server := server.NewServer(cfg, r)
	server.Run()
}
