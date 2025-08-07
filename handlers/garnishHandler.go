package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/iglesiasgm/wiki-cocktail/db"
	"github.com/iglesiasgm/wiki-cocktail/models"
)

func GetAllGarnishes(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT idgarnish, name FROM garnish")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var garnishes []models.Garnish
	for rows.Next() {
		var g models.Garnish
		if err := rows.Scan(&g.ID, &g.Name); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		garnishes = append(garnishes, g)
	}

	// Devolver JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(garnishes)
}
