package main

import (
	"fmt"
	"log"
	"net/http"

	"api_raz_mat/internal/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Inicializa el enrutador
	r := mux.NewRouter()

	// Agrega algunas rutas de ejemplo
	r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/books", handlers.AddBook).Methods("POST")

	// Configura el servidor web
	fmt.Println("Servidor iniciado en http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
