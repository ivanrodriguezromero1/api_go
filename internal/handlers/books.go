package handlers

import (
	"api_raz_mat/internal/models"
	"encoding/json"
	"net/http"
)

var books []models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	if err := validateAPIKey(r); err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
func AddBook(w http.ResponseWriter, r *http.Request) {
	if err := validateAPIKey(r); err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}
func validateAPIKey(r *http.Request) error {
	// apiKey := r.Header.Get("X-API-Key")
	// if apiKey != "my-secret-api-key" {
	// 	return errors.New("API key missing or invalid")
	// }
	return nil
}
