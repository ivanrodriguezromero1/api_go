package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Problem struct {
	Difficulty string `json:"difficulty"`
	Language   string `json:"language"`
	Sequence   string `json:"sequence"`
	Answer     int    `json:"answer"`
	Options    []int  `json:"options"`
}

func main() {
	http.HandleFunc("/problem", problemHandler)
	http.ListenAndServe(":8000", nil)
}

func problemHandler(w http.ResponseWriter, r *http.Request) {
	// Parseamos los parámetros
	dificultad := r.URL.Query().Get("dificultad")
	idioma := r.URL.Query().Get("idioma")

	// Calculamos el problema
	secuencia := getSecuencia(dificultad, idioma)
	respuesta := calculateAnswer(secuencia)
	opciones := generateOptions(respuesta)
	serie := fmt.Sprintf("%v", secuencia)
	// Creamos el problema
	problem := Problem{
		Difficulty: dificultad,
		Language:   idioma,
		Sequence:   serie,
		Answer:     respuesta,
		Options:    opciones,
	}
	// Devolvemos la respuesta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(problem)
}

// Retorna una serie aritmética según la dificultad e idioma ingresado
func getSecuencia(dificultad, idioma string) []int {
	// Define las series aritméticas según la dificultad e idioma
	series := map[string]map[string][]int{
		"facil": {
			"es": []int{1, 2, 3, 4, 5},
			"en": []int{2, 4, 6, 8, 10},
		},
		"medio": {
			"es": []int{3, 6, 9, 12, 15},
			"en": []int{1, 3, 5, 7, 9},
		},
		"dificil": {
			"es": []int{1, 4, 7, 10, 13},
			"en": []int{2, 5, 8, 11, 14},
		},
	}
	// Retorna la serie aritmética correspondiente
	return series[dificultad][idioma]
	// return fmt.Sprintf("%v", series[dificultad][idioma])
}

// Calcula la respuesta del problema
func calculateAnswer(s []int) int {
	// var s []int
	// json.Unmarshal([]byte(secuencia), &s)
	// fmt.Println(secuencia)
	// fmt.Println(s)
	// Calcula la respuesta
	return (s[len(s)-1] + s[0]) * len(s) / 2
}

// Genera tres opciones aleatorias para despistar al usuario
func generateOptions(answer int) []int {
	rand.Seed(time.Now().UnixNano())
	options := make([]int, 4)
	for i := range options {
		options[i] = answer + rand.Intn(20) - 10
	}
	options[rand.Intn(4)] = answer
	return options
}
