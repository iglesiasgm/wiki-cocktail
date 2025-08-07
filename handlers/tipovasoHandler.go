package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/iglesiasgm/wiki-cocktail/db"
	"github.com/iglesiasgm/wiki-cocktail/models"
)

func GetAllTipoVaso(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT idTipoVaso, nombreTipoVaso FROM tipoVaso")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tipoVaso []models.TipoVaso
	for rows.Next() {
		var g models.TipoVaso
		if err := rows.Scan(&g.ID, &g.NOMBRE); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tipoVaso = append(tipoVaso, g)
	}

	// Devolver JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tipoVaso)
}
