package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/iglesiasgm/wiki-cocktail/db"
	"github.com/iglesiasgm/wiki-cocktail/models"
)

func GetAllTipoHielo(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT idtipohielo, nombretipohielo FROM tipohielo")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tipoHielo []models.TipoHielo
	for rows.Next() {
		var g models.TipoHielo
		if err := rows.Scan(&g.ID, &g.NOMBRE); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tipoHielo = append(tipoHielo, g)
	}

	// Devolver JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tipoHielo)
}
