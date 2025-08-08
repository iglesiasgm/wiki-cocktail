package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/iglesiasgm/wiki-cocktail/db"
	"github.com/iglesiasgm/wiki-cocktail/models"
)

func GetAllTipoIngrediente(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT codtipoingrediente, nombreti FROM tipoingrediente")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tipoingrediente []models.TipoIngrediente
	for rows.Next() {
		var g models.TipoIngrediente
		if err := rows.Scan(&g.COD, &g.Nombre, &g.UnidadMedida); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tipoingrediente = append(tipoingrediente, g)
	}

	// Devolver JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tipoingrediente)
}
