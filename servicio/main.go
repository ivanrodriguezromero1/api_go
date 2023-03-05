package main

import (
	"api_raz_mat/internal/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Inicializa el enrutador
	r := mux.NewRouter()
	// Agrega algunas rutas de ejemplo
	r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/books", handlers.AddBook).Methods("POST")

	// Configura el servidor web
	fmt.Println("Servidor iniciado")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
}
