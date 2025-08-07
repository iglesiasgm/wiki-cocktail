package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/iglesiasgm/wiki-cocktail/db"
	"github.com/iglesiasgm/wiki-cocktail/models"
)

func GetAllUnidadMedida(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT codUM, unidadMedida, cantidadIC FROM unidadMedida")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var unidadMedida []models.UnidadMedida
	for rows.Next() {
		var g models.UnidadMedida
		if err := rows.Scan(&g.COD, &g.UM, &g.CANTIDAD); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		unidadMedida = append(unidadMedida, g)
	}

	// Devolver JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(unidadMedida)
}
