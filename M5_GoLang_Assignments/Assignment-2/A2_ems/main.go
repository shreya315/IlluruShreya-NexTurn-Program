package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// Assuming Product struct is imported from model.go
// If main.go and model.go are in the same package, no import is needed
// Otherwise, use: "your_module_name/model" to import the Product struct from model.go

// Global variable to simulate database
var products []Product
var nextID = 1

// Create a new product (POST /products)
func createProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	newProduct.ID = nextID
	nextID++
	products = append(products, newProduct)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newProduct)
}

// Get all products (GET /products)
func getAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// Get a single product by ID (GET /products/{id})
func getProductByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/products/"):]
	productID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	for _, product := range products {
		if product.ID == productID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(product)
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

// Update a product by ID (PUT /products/{id})
func updateProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/products/"):]
	productID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var updatedProduct Product
	if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	for i, product := range products {
		if product.ID == productID {
			updatedProduct.ID = product.ID
			products[i] = updatedProduct
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedProduct)
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

// Delete a product by ID (DELETE /products/{id})
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/products/"):]
	productID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	for i, product := range products {
		if product.ID == productID {
			products = append(products[:i], products[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Product deleted successfully"})
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

func main() {
	// Handle routes
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			createProduct(w, r)
		} else if r.Method == "GET" {
			getAllProducts(w, r)
		}
	})

	http.HandleFunc("/products/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getProductByID(w, r)
		} else if r.Method == "PUT" {
			updateProduct(w, r)
		} else if r.Method == "DELETE" {
			deleteProduct(w, r)
		}
	})

	log.Println("Server started on http://localhost:8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}
