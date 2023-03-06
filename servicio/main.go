package main

import (
	"api_raz_mat/internal/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Inicializa el enrutador
	r := mux.NewRouter()
	// Agrega algunas rutas de ejemplo
	r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/books", handlers.AddBook).Methods("POST")
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization"},
	})
	handler := c.Handler(r)
	fmt.Println("Servidor iniciado")
	if err := http.ListenAndServe(":8000", handler); err != nil {
		log.Fatal(err)
	}
}
