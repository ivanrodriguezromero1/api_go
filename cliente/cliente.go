package main

import (
	"api_raz_mat/internal/models"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

func main() {
	const apiKey = "my-secret-api-key"
	url := "https://api-go2.onrender.com/books"
	bulk := flag.Bool("bulk", false, "Insertar datos por lotes")
	insert := flag.Bool("insert", false, "Insertar nuevo dato")
	show := flag.Bool("show", false, "Mostrar datos")
	flag.Parse()
	if *bulk {
		// Agrega libros
		books := []models.Book{
			{Title: "El nombre del viento", Author: "Patrick Rothfuss"},
			{Title: "El temor de un hombre sabio", Author: "Patrick Rothfuss"},
			{Title: "Cien años de soledad", Author: "Gabriel García Márquez"},
		}
		for _, book := range books {
			addBook(book, url, apiKey)
		}
	}
	if *insert {
		// Agrega un nuevo libro
		newBook := models.Book{
			Title:  "La sombra del viento",
			Author: "Carlos Ruiz Zafón",
		}
		addBook(newBook, url, apiKey)
	}
	if *show {
		// Obtiene todos los libros
		getBooks(url, apiKey)
	}
}

func addBook(book models.Book, url, apiKey string) {
	bookJson, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bookJson))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("X-API-Key", apiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	var createdBook models.Book
	err = json.NewDecoder(resp.Body).Decode(&createdBook)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Se creó el libro: %+v\n", createdBook)
}

func getBooks(url string, apiKey string) {
	// Crea la solicitud HTTP
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Agrega el encabezado personalizado con la API key
	req.Header.Set("X-API-Key", apiKey)

	// Envía la solicitud HTTP y maneja la respuesta
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	var books []models.Book
	err = json.NewDecoder(resp.Body).Decode(&books)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Libros:")
	for _, book := range books {
		fmt.Printf("- %+v\n", book)
	}
}
