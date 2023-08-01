package main

import (
	"net/http"

	"github.com/axelldev/products-api/server"
	"github.com/gorilla/mux"
)

func main() {
	cfg := server.NewServerConfig()

	r := mux.NewRouter()
	r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		Json(w, http.StatusOK, map[string]bool{"ok": true})
	})

	server := server.NewServer(cfg, r)
	server.Run()
}
