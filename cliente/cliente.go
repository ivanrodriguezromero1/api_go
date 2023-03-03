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
			addBook(book)
		}
	}
	if *insert {
		// Agrega un nuevo libro
		newBook := models.Book{
			Title:  "La sombra del viento",
			Author: "Carlos Ruiz Zafón",
		}
		addBook(newBook)
	}
	if *show {
		// Obtiene todos los libros
		getBooks()
	}
}

func addBook(book models.Book) {
	bookJson, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
		return
	}

	req, err := http.NewRequest("POST", "http://localhost:8000/books", bytes.NewBuffer(bookJson))
	if err != nil {
		fmt.Println(err)
		return
	}

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

func getBooks() {
	resp, err := http.Get("http://localhost:8000/books")
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
