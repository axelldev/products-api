package handler

import (
	"net/http"

	"github.com/axelldev/products-api/models"
	"github.com/axelldev/products-api/response"
)

// HandleProducts handles all requests to /api/products
func HandleProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		listProducts(w, r)
	}

	if r.Method == http.MethodPost {

	}

	if r.Method == http.MethodPut {

	}

	if r.Method == http.MethodDelete {
	}
}

// Endpoint: GET /api/products
// List all products
func listProducts(w http.ResponseWriter, r *http.Request) {
	products := []models.Product{
		{
			ID:          1,
			Name:        "Product 1",
			Description: "Description 1",
			Price:       10.50,
		},
	}

	response.Json(w, http.StatusOK, products)
}
