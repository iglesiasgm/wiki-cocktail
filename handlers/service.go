package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/iglesiasgm/wiki-cocktail/data"
	"github.com/iglesiasgm/wiki-cocktail/models"
)

func GetCocktails(w http.ResponseWriter, r *http.Request) {
	log.Println("Ruta recibida:", r.URL.Path, "Método:", r.Method)
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data.Cocktails)
	case http.MethodPost:
		CreateCocktail(w, r)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func GetCocktailByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// /cocktails/2
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.Error(w, "Ruta no válida", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for _, c := range data.Cocktails {
		if c.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(c)
			return
		}
	}

	http.Error(w, "Cóctel no encontrado", http.StatusNotFound)
}

func CreateCocktail(w http.ResponseWriter, r *http.Request) {
	var newCocktail models.Cocktail
	err := json.NewDecoder(r.Body).Decode(&newCocktail)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	newCocktail.ID = len(data.Cocktails) + 1
	data.Cocktails = append(data.Cocktails, newCocktail)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCocktail)
}
