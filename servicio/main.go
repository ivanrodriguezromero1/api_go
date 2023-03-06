package main

import (
	"fmt"
	"net/http"
)

func main() {

	// r := mux.NewRouter()
	// r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	// r.HandleFunc("/books", handlers.AddBook).Methods("POST")
	// c := cors.New(cors.Options{
	// 	AllowedOrigins: []string{"*"},
	// 	AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowedHeaders: []string{"Accept", "Content-Length", "Accept-Encoding"},
	// })
	// handler := c.Handler(r)
	// fmt.Println("Servidor iniciado")
	// if err := http.ListenAndServe(":8000", handler); err != nil {
	// 	log.Fatal(err)
	// }

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Permitimos todas las solicitudes de cualquier origen
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if r.Method == "OPTIONS" {
			// Agregamos los headers CORS necesarios para las solicitudes OPTIONS
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.WriteHeader(http.StatusOK)
			return
		}

		// Enviamos la respuesta normal
		fmt.Fprint(w, "Hola, mundo!")
	})

	fmt.Println("Servidor iniciado")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
