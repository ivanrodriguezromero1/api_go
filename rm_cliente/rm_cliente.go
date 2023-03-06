package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type Problem struct {
	Difficulty string `json:"difficulty"`
	Language   string `json:"language"`
	Sequence   string `json:"sequence"`
	Answer     int    `json:"answer"`
	Options    []int  `json:"options"`
}

func main() {
	apiUrl := "http://localhost:8000/problem"
	params := url.Values{}
	params.Add("dificultad", "facil")
	params.Add("idioma", "es")
	url := apiUrl + "?" + params.Encode()

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var problem Problem
	err = json.NewDecoder(response.Body).Decode(&problem)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Difficulty:", problem.Difficulty)
	fmt.Println("Language:", problem.Language)
	fmt.Println("Sequence:", problem.Sequence)
	fmt.Println("Answer:", problem.Answer)
	fmt.Println("Options:", problem.Options)
}
