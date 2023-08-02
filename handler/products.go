package handler

import (
	"net/http"

	"github.com/axelldev/products-api/models"
	"github.com/axelldev/products-api/response"
	"github.com/gorilla/mux"
)

// HandleProducts handles all requests to /api/products
func HandleProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		listProducts(w, r)
	}

	if r.Method == http.MethodPost {
		createProduct(w, r)
	}

	if r.Method == http.MethodPut {
		updateProduct(w, r)
	}

	if r.Method == http.MethodDelete {
		updateProduct(w, r)
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

func DetailProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)
	response.Json(w, http.StatusOK, id)
}

// Endpoint: POST /api/products
// Create a new product
func createProduct(w http.ResponseWriter, r *http.Request) {
	response.Json(w, http.StatusOK, nil)
}

// Endpoint: PUT /api/products/{id}
// Update a product
func updateProduct(w http.ResponseWriter, r *http.Request) {
	response.Json(w, http.StatusOK, nil)
}

// Endpoint: DELETE /api/products/{id}
// Delete a product
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	response.Json(w, http.StatusOK, nil)
}
