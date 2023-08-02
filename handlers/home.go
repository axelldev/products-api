package handlers

import (
	"net/http"

	"github.com/axelldev/products-api/response"
	"github.com/axelldev/products-api/server"
)

type HomeResponse struct {
	Message string `json:"message"`
	OK      bool   `json:"ok"`
}

func HomeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.Json(w, http.StatusOK, HomeResponse{
			Message: "Server is running!",
			OK:      true,
		})
	}
}
