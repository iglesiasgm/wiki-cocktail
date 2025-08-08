package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/iglesiasgm/wiki-cocktail/db"
	"github.com/iglesiasgm/wiki-cocktail/models"
)

func GetAllMetodo(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT idmetodo, descripmetodo FROM metodo")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var metodos []models.Metodo
	for rows.Next() {
		var g models.Metodo
		if err := rows.Scan(&g.ID, &g.Description); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		metodos = append(metodos, g)
	}

	// Devolver JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metodos)
}
