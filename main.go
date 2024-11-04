package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var products = []Product{
	{ID: 1, Name: "Product 1"},
	{ID: 2, Name: "Product 2"},
}

func getAllProducts(w http.ResponseWriter, r *http.Request) {
	// Возвращает все продукты
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	// Обновляет товар
	fmt.Fprint(w, "Product updated")
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	// Удатяет продукт
	fmt.Fprint(w, "Product deleted")
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	// Добавляет новый продукт
	fmt.Fprint(w, "Product created")
}

func getProductByID(w http.ResponseWriter, r *http.Request) {
	// Возвращает конкретный продукт по id
	fmt.Fprint(w, "Product details")
}

func main() {
	http.HandleFunc("/products", getAllProducts)
	http.HandleFunc("/products/update", updateProduct)
	http.HandleFunc("/products/delete", deleteProduct)
	http.HandleFunc("/products/create", createProduct)
	http.HandleFunc("/products/{id}", getProductByID)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
